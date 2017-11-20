/*
 * Copyright (c) 2015, Shinya Yagyu
 * All rights reserved.
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 * 3. Neither the name of the copyright holder nor the names of its
 *    contributors may be used to endorse or promote products derived from this
 *    software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 * See LICENSE file for the original license:
 */

package gocoin

const (
	op0                   = byte(0)
	opFALSE               = byte(0)
	opNA                  = byte(1)
	opPUSHDATA1           = byte(76)
	opPUSHDATA2           = byte(77)
	opPUSHDATA4           = byte(78)
	op1NEGATE             = byte(79)
	opTRUE                = byte(81)
	opNOP                 = byte(97)
	opIF                  = byte(99)
	opNOTIF               = byte(100)
	opELSE                = byte(103)
	opENDIF               = byte(104)
	opVERIFY              = byte(105)
	opRETURN              = byte(106)
	opTOALTSTACK          = byte(107)
	opFROMALTSTACK        = byte(108)
	opIFDUP               = byte(115)
	opDEPTH               = byte(116)
	opDROP                = byte(117)
	opDUP                 = byte(118)
	opNIP                 = byte(119)
	opOVER                = byte(120)
	opPICK                = byte(121)
	opROLL                = byte(122)
	opROT                 = byte(123)
	opSWAP                = byte(124)
	opTUCK                = byte(125)
	op2DROP               = byte(109)
	op2DUP                = byte(110)
	op3DUP                = byte(111)
	op2OVER               = byte(112)
	op2ROT                = byte(113)
	op2SWAP               = byte(114)
	opCAT                 = byte(126)
	opSUBSTR              = byte(127)
	opLEFT                = byte(128)
	opRIGHT               = byte(129)
	opSIZE                = byte(130)
	opINVERT              = byte(131)
	opAND                 = byte(132)
	opOR                  = byte(133)
	opXOR                 = byte(134)
	opEQUAL               = byte(135)
	opEQUALVERIFY         = byte(136)
	op1ADD                = byte(139)
	op1SUB                = byte(140)
	op2MUL                = byte(141)
	op2DIV                = byte(142)
	opNEGATE              = byte(143)
	opABS                 = byte(144)
	opNOT                 = byte(145)
	op0NOTEQUAL           = byte(146)
	opADD                 = byte(147)
	opSUB                 = byte(148)
	opMUL                 = byte(149)
	opDIV                 = byte(150)
	opMOD                 = byte(151)
	opLSHIFT              = byte(152)
	opRSHIFT              = byte(153)
	opBOOLAND             = byte(154)
	opBOOLOR              = byte(155)
	opNUMEQUAL            = byte(156)
	opNUMEQUALVERIFY      = byte(157)
	opNUMNOTEQUAL         = byte(158)
	opLESSTHAN            = byte(159)
	opGREATERTHAN         = byte(160)
	opLESSTHANOREQUAL     = byte(161)
	opGREATERTHANOREQUAL  = byte(162)
	opMIN                 = byte(163)
	opMAX                 = byte(164)
	opWITHIN              = byte(165)
	opRIPEMD160           = byte(166)
	opSHA1                = byte(167)
	opSHA256              = byte(168)
	opHASH160             = byte(169)
	opHASH256             = byte(170)
	opCODESEPARATOR       = byte(171)
	opCHECKSIG            = byte(172)
	opCHECKSIGVERIFY      = byte(173)
	opCHECKMULTISIG       = byte(174)
	opCHECKMULTISIGVERIFY = byte(175)
	opPUBKEYHASH          = byte(253)
	opPUBKEY              = byte(254)
	opINVALIDOPCODE       = byte(255)
	opRESERVED            = byte(80)
	opVER                 = byte(98)
	opVERIF               = byte(101)
	opVERNOTIF            = byte(102)
	opRESERVED1           = byte(137)
	opRESERVED2           = byte(138)
	opNOP1                = byte(176)
	opNOP2                = byte(177)
	opNOP3                = byte(178)
	opNOP4                = byte(179)
	opNOP5                = byte(180)
	opNOP6                = byte(181)
	opNOP7                = byte(182)
	opNOP8                = byte(183)
	opNOP9                = byte(184)
	opNOP10               = byte(185)
	op1                   = byte(81)
	op2                   = byte(82)
	op3                   = byte(83)
	op4                   = byte(84)
	op5                   = byte(85)
	op6                   = byte(86)
	op7                   = byte(87)
	op8                   = byte(88)
	op9                   = byte(89)
	op10                  = byte(90)
	op11                  = byte(91)
	op12                  = byte(92)
	op13                  = byte(93)
	op14                  = byte(94)
	op15                  = byte(95)
	op16                  = byte(96)

	//BTC is unit to convert BTC to satoshi
	BTC = 100000000

	//Fee for a transaction
	DefaultFee = uint64(0.0001 * BTC) //  0.0001 BTC/kB
)
