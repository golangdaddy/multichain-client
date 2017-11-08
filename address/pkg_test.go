package address

import (
    "fmt"
    "testing"
)

func TestAddress(t *testing.T) {

    t.Run("Test address generation", func (t *testing.T) {

        address, err := MultiChain(fmt.Sprintf("0284E5235E299AF81EBE1653AC5F06B60E13A3A81F918018CBD10CE695095B3E24"))

        if err != nil {
            t.Error(err)
        }

        if address != "1Yu2BuptuZSiBWfr2Qy4aic6qEVnwPWrdkHPEc" {
            t.Error("INVALID PUBLIC ADDRESSS GENERATED")
        }

    })

}
