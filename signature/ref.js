'use strict';

var randomBytes = require('randombytes'),
    logger = require('./logger'),
    secp256k1 = require('secp256k1'),
    crypto = require('./crypto'),
    base58 = require('bs58'),
    xor = require('buffer-xor'),
    compare = require('buffer-compare'),
    Promise = require('bluebird'),
    clone = require('clone');

var OPS = require('./opcodes');

var _generatePrivateKey, _createPublicKey, _generateAddress, _extendWithVersion, _generateChecksum, _extractVersion, _toBuffer;

var addressFactory = {};

addressFactory.generateNew = function(pubKeyHashVersion, checksumValue, randomBytesGenerator) {
    return _generateAddress(_generatePrivateKey(randomBytesGenerator), pubKeyHashVersion, checksumValue, false);
};

addressFactory.fromWIF = function(wif, publicKeyHashVersion, privateKeyVersion, checksumValue) {
    privateKeyVersion = Buffer.isBuffer(privateKeyVersion) ? privateKeyVersion : Buffer.from(privateKeyVersion, 'hex');
    checksumValue = Buffer.isBuffer(checksumValue) ? checksumValue : Buffer.from(checksumValue, 'hex');

    var decodedWIF = new Buffer(base58.decode(wif));
    logger.log('[from WIF]', 'decoded WIF', decodedWIF.toString('hex'));

    var extractedChecksum = decodedWIF.slice(decodedWIF.length - checksumValue.length),
        extendedPrivateKey = decodedWIF.slice(0, decodedWIF.length - checksumValue.length),
        generatedChecksum = _generateChecksum(extendedPrivateKey, checksumValue.length),
        xorChecksum = xor(generatedChecksum, checksumValue);

    logger.log('[from WIF]', 'extracted checksum', extractedChecksum.toString('hex'));
    logger.log('[from WIF]', 'extended private key', extendedPrivateKey.toString('hex'));

    logger.log('[from WIF]', 'generated checksum', generatedChecksum.toString('hex'));
    logger.log('[from WIF]', 'xor checksum', xorChecksum.toString('hex'));

    if (compare(extractedChecksum, xorChecksum) !== 0) {
        throw new Error('Extracted checksum and generated checksum do not match (' + extractedChecksum.toString('hex') + ', ' + xorChecksum.toString('hex') + ')');
    }

    var extractedData = _extractVersion(extendedPrivateKey, privateKeyVersion.length, 8);

    if (compare(extractedData['version'], privateKeyVersion) !== 0) {
        throw new Error('Extracted private key does not match the given private key (' + extractedData['version'].toString('hex') + ', ' + privateKeyVersion.toString('hex') + ')');
    }

    var privateKey = extractedData['hash'],
        compressed = false;

    logger.log('[from WIF]', 'extracted private key', privateKey.toString('hex'));

    if (privateKey.length !== 32) {
        if (privateKey.length === 33 && privateKey[32] === 1) {
            compressed = true;
            privateKey = privateKey.slice(0, 32);
        } else {
            throw new Error('Private key length invalid ' + privateKey.length + ' bytes');
        }
    }

    return _generateAddress(privateKey, publicKeyHashVersion, checksumValue, compressed);
};

module.exports = addressFactory;

function Address(encodedAddress, pubKey, privKey, compressed) {
    this.address = encodedAddress;
    this.publicKey = pubKey;
    this.privateKey = privKey;
    this.compressed = compressed
}

Address.prototype.toString = function() {
    return this.address + ' (' + this.privateKey.toString('hex') + ')';
};

Address.prototype.sign = function(rawTransaction, index, getInputScript) {
    rawTransaction = clone(rawTransaction);
    index = index || 0;

    logger.log('[sign]', 'unsigned to buffer', _toBuffer(rawTransaction).toString('hex'));

    var scriptPromise;
    switch(typeof getInputScript) {
        case 'function':
            scriptPromise = getInputScript(rawTransaction.vin[index]['txid'], rawTransaction.vin[index]['vout']);
            break;
        case 'object':
            scriptPromise = getInputScript;
            break;
        case 'string':
            scriptPromise = Promise.resolve(getInputScript);
            break;
    }

    var self = this;

    return scriptPromise.then(function(script) {
        rawTransaction.vin[index].script = Buffer.isBuffer(script) ? script : Buffer.from(script, 'hex');

        logger.log('[sign]', 'script', script.toString('hex'));

        var hashType = 0x01;  // SIGHASH_ALL
        var hashForSignature = crypto.hash256(Buffer.concat([_toBuffer(rawTransaction), uint32Buffer(hashType)]));

        logger.log('[sign]', 'hash for signature', hashForSignature.toString('hex'));

        var signature = secp256k1.sign(hashForSignature, self.privateKey).signature;
        var signatureDER = secp256k1.signatureExport(signature);

        logger.log('[sign]', 'signature', signature.toString('hex'));
        logger.log('[sign]', 'signature DER', signatureDER.toString('hex'));

        var scriptSignature = Buffer.concat([signatureDER, uint8Buffer(hashType)]); // public key hash input

        logger.log('[sign]', 'script signature', scriptSignature.toString('hex'));

        var scriptSig = Buffer.concat([pushDataIntBuffer(scriptSignature.length), scriptSignature, pushDataIntBuffer(self.publicKey.length), self.publicKey]);

        logger.log('[sign]', 'script sig', scriptSig.toString('hex'));

        rawTransaction.vin[index].script = scriptSig;

        rawTransaction.toBuffer = function() {

            return _toBuffer(this);
        }

        return rawTransaction;
    });
};

// @todo better way of doing this ... wrapping it in an object? move to utils?
Address.prototype.toBuffer = function(rawTransaction) {

    return _toBuffer(rawTransaction);
}

Address.prototype.toWIF = function(privateKeyVersion, checksumValue) {
    privateKeyVersion = Buffer.isBuffer(privateKeyVersion) ? privateKeyVersion : Buffer.from(privateKeyVersion, 'hex');
    checksumValue = Buffer.isBuffer(checksumValue) ? checksumValue : Buffer.from(checksumValue, 'hex');

    var privateKey = this.privateKey;

    logger.log('[to WIF]', 'private key', privateKey.toString('hex'));

    if (this.compressed) {
        privateKey = Buffer.concat(privateKey, Buffer.from('01', 'hex'));
        logger.log('[to WIF]', 'add compressed flag', privateKey.toString('hex'));
    }

    var extendedPrivateKey = _extendWithVersion(privateKey, privateKeyVersion, 8);

    logger.log('[to WIF]', 'extended private key', extendedPrivateKey.toString('hex'));

    var checksum = _generateChecksum(extendedPrivateKey, checksumValue.length),
        xorChecksum = xor(checksum, checksumValue);

    logger.log('[to WIF]', 'checksum', checksum.toString('hex'));
    logger.log('[to WIF]', 'xor checksum', xorChecksum.toString('hex'));

    var decodedWIF = Buffer.concat([extendedPrivateKey, xorChecksum]);
    logger.log('[to WIF]', 'decoded WIF', decodedWIF.toString('hex'));

    var encodedWIF = base58.encode(decodedWIF);

    logger.log('[to WIF]', 'encoded WIF', encodedWIF);

    return encodedWIF;
};

_generatePrivateKey = function(randomBytesGenerator) {
    var privKey;

    randomBytesGenerator = randomBytesGenerator && typeof randomBytesGenerator === 'function' ? randomBytesGenerator : randomBytes;

    do {
        privKey = randomBytesGenerator(32);
        if (!Buffer.isBuffer(privKey)) {
            if (typeof privKey === 'string') {
                privKey = Buffer.from(privKey, 'hex');
            } else {
                throw new Error('Invalid random bytes generator');
            }
        }
    } while (!secp256k1.privateKeyVerify(privKey));

    return privKey;
};

_generateAddress = function(privKey, pubKeyHashVersion, checksumValue, compressed) {
    pubKeyHashVersion = Buffer.isBuffer(pubKeyHashVersion) ? pubKeyHashVersion : Buffer.from(pubKeyHashVersion, 'hex');
    checksumValue = Buffer.isBuffer(checksumValue) ? checksumValue : Buffer.from(checksumValue, 'hex');

    logger.log('[Generate address]', 'private key', privKey.toString('hex'));

    var pubKey = _createPublicKey(privKey, compressed);

    logger.log('[Generate address]', 'public key', pubKey.toString('hex'));

    var ripemd160 = crypto.ripemd160(crypto.sha256(pubKey)),
        extendedRipemd160 = _extendWithVersion(ripemd160, pubKeyHashVersion, 5);

    logger.log('[Generate address]', 'ripemd160', ripemd160.toString('hex'));
    logger.log('[Generate address]', 'public key hash value', pubKeyHashVersion.toString('hex'));
    logger.log('[Generate address]', 'extended ripemd160', extendedRipemd160.toString('hex'));

    var checksum = _generateChecksum(extendedRipemd160, checksumValue.length),
        xorChecksum = xor(checksum, checksumValue);
    logger.log('[Generate address]', 'checksum', checksum.toString('hex'));
    logger.log('[Generate address]', 'xor checksum', xorChecksum.toString('hex'));

    var decodedAddress = Buffer.concat([extendedRipemd160, xorChecksum]);
    logger.log('[Generate address]', 'decoded address', decodedAddress.toString('hex'));

    var encodedAddress = base58.encode(decodedAddress);

    logger.log('[Generate address]', 'encoded address', encodedAddress);

    return new Address(encodedAddress, pubKey, privKey, compressed);
};

_createPublicKey = function(privKey, combined) {
    combined = combined || false;

    return secp256k1.publicKeyCreate(privKey, combined);
};

_extendWithVersion = function(hash, versionHash, nbSpacerBytes) {
    var extendedParts = [], index = 0, fromIndex, toIndex;

    for (; index < versionHash.length; index++) {
        extendedParts.push(versionHash.slice(index, index + 1));

        fromIndex = index * nbSpacerBytes;
        toIndex = (index + 1) * nbSpacerBytes;

        extendedParts.push(hash.slice(fromIndex, toIndex));
    }

    if ((index * nbSpacerBytes) < hash.length) {
        extendedParts.push(hash.slice(index * nbSpacerBytes));
    }

    return Buffer.concat(extendedParts);
};

_extractVersion = function(extendedHash, versionLength, nbSpacerBytes) {
    var versionParts = [],
        hashParts = [], index = 0, fromIndex, toIndex;

    for (; index < versionLength; index++) {
        versionParts.push(extendedHash.slice(index * nbSpacerBytes + index, index * nbSpacerBytes + index + 1));

        fromIndex = index * nbSpacerBytes + index + 1;
        toIndex = (index + 1) * nbSpacerBytes + index + 1;

        hashParts.push(extendedHash.slice(fromIndex, toIndex));
    }

    if ((index * nbSpacerBytes + index) < extendedHash.length) {
        hashParts.push(extendedHash.slice(index * nbSpacerBytes + index));
    }

    return {
        'version': Buffer.concat(versionParts),
        'hash': Buffer.concat(hashParts)
    };
};

_generateChecksum = function(extendedHash, checksumLength) {
    return crypto.hash256(extendedHash).slice(0, checksumLength);
};


_toBuffer = function(decodedTransaction) {
    var chunks = [];

    chunks.push(uint32Buffer(decodedTransaction.version));
    chunks.push(varIntBuffer(decodedTransaction.vin.length));

    decodedTransaction.vin.forEach(function (txIn, index) {
        var hash = [].reverse.call(new Buffer(txIn.txid, 'hex'));
        chunks.push(hash);
        chunks.push(uint32Buffer(txIn.vout)); // index

        if (txIn.script != null) {
            logger.log('[toBuffer]', Buffer.concat(chunks).toString('hex'));
            logger.log('[toBuffer]', 'length', txIn.script.length);
            chunks.push(varIntBuffer(txIn.script.length));
            logger.log('[toBuffer]', 'script', txIn.script.toString('hex'));
            chunks.push(txIn.script);
        } else {
            chunks.push(varIntBuffer(0));
        }

        chunks.push(uint32Buffer(txIn.sequence));
    });

    chunks.push(varIntBuffer(decodedTransaction.vout.length));
    decodedTransaction.vout.forEach(function (txOut) {
        chunks.push(uint64Buffer(txOut.value));

        var script = Buffer.from(txOut.scriptPubKey.hex, 'hex');

        chunks.push(varIntBuffer(script.length));
        chunks.push(script);
    });

    chunks.push(uint32Buffer(decodedTransaction.locktime));

    return Buffer.concat(chunks);
};

var pushDataIntBuffer = function(number) {
    var chunks = [];

    var pushDataSize = number < OPS.OP_PUSHDATA1 ? 1
        : number < 0xff ? 2
        : number < 0xffff ? 3
        : 5;

    if (pushDataSize === 1) {
        chunks.push(uint8Buffer(number));
    } else if (pushDataSize === 2) {
        chunks.push(uint8Buffer(OPS.OP_PUSHDATA1));
        chunks.push(uint8Buffer(number));
    } else if (pushDataSize === 3) {
        chunks.push(uint8Buffer(OPS.OP_PUSHDATA2));
        chunks.push(uint16Buffer(number));
    } else {
        chunks.push(uint8Buffer(OPS.OP_PUSHDATA4));
        chunks.push(uint32Buffer(number));
    }

    return Buffer.concat(chunks);
};

var varIntBuffer = function(number) {
    var chunks = [];

    var size = number < 253 ? 1
        : number < 0x10000 ? 3
        : number < 0x100000000 ? 5
        : 9;

    // 8 bit
    if (size === 1) {
        chunks.push(uint8Buffer(number));

        // 16 bit
    } else if (size === 3) {
        chunks.push(uint8Buffer(253));
        chunks.push(uint16Buffer(number));

        // 32 bit
    } else if (size === 5) {
        chunks.push(uint8Buffer(254));
        chunks.push(uint32Buffer(number));

        // 64 bit
    } else {
        chunks.push(uint8Buffer(255));
        chunks.push(uint64Buffer(number));
    }

    return Buffer.concat(chunks);
};

var uint8Buffer = function(number) {
    var buffer = new Buffer(1);
    buffer.writeUInt8(number, 0);

    return buffer;
};

var uint16Buffer = function(number) {
    var buffer = new Buffer(2);
    buffer.writeUInt16LE(number, 0);

    return buffer;
};

var uint32Buffer = function(number) {
    var buffer = new Buffer(4);
    buffer.writeUInt32LE(number, 0);

    return buffer;
};

var uint64Buffer = function(number) {
    var buffer = new Buffer(8);
    buffer.writeInt32LE(number & -1, 0)
    buffer.writeUInt32LE(Math.floor(number / 0x100000000), 4)

    return buffer;
};