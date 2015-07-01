//Semtech SX1509

package sx1509

import "github.com/NeuralSpaz/i2c"

const (
	SX1509_ADDRESS01            = 0x3E
	SX1509_ADDRESS02            = 0x3F
	SX1509_REG_INPUTDISABLEB    = 0x00
	SX1509_REG_INPUTDISABLEA    = 0x01
	SX1509_REG_LONGSLEWB        = 0x02
	SX1509_REG_LONGSLEWA        = 0x03
	SX1509_REG_LOWDRIVEB        = 0x04
	SX1509_REG_LOWDRIVEA        = 0x05
	SX1509_REG_PULLUPB          = 0x06
	SX1509_REG_PULLUPA          = 0x07
	SX1509_REG_PULLDOWNB        = 0x08
	SX1509_REG_PULLDOWNA        = 0x09
	SX1509_REG_OPENDRAINB       = 0x0A
	SX1509_REG_OPENDRAINA       = 0x0B
	SX1509_REG_POLARITYB        = 0x0C
	SX1509_REG_POLARITYA        = 0x0D
	SX1509_REG_DIRB             = 0x0E
	SX1509_REG_DIRA             = 0x0F
	SX1509_REG_DATAB            = 0x10
	SX1509_REG_DATAA            = 0x11
	SX1509_REG_INTERRUPTMASKB   = 0x12
	SX1509_REG_INTERRUPTMASKA   = 0x13
	SX1509_REG_SENSEHIGHB       = 0x14
	SX1509_REG_SENSELOWB        = 0x15
	SX1509_REG_SENSEHIGHA       = 0x16
	SX1509_REG_SENSELOWA        = 0x17
	SX1509_REG_INTERRUPTSOURCEB = 0x18
	SX1509_REG_INTERRUPTSOURCEA = 0x19
	SX1509_REG_EVENTSTATUSB     = 0x1A
	SX1509_REG_EVENTSTATUSA     = 0x1B
	SX1509_REG_LEVELSHIFTER1    = 0x1C
	SX1509_REG_LEVELSHIFTER2    = 0x1D
	SX1509_REG_CLOCK            = 0x1E
	SX1509_REG_MISC             = 0x1F
	SX1509_REG_LEDDRIVERENABLEB = 0x20
	SX1509_REG_LEDDRIVERENABLEA = 0x21
)

type SX1509 struct {
	Dev         i2c.I2CBus
	initialized bool
	Address     uint8
}

func New(deviceAdress uint8, i2cbus byte) *SX1509 {
	deviceBus := i2c.NewI2CBus(i2cbus)
	d := &SX1509{
		Dev:     deviceBus,
		Address: deviceAdress,
	}
	return d
}


func (d *SX1509) WritePORT(port uint8, data byte) error {
	if !(d.initialized) {
		if err := initSX1509(d); err != nil {
			return err
		}
	}
	if port == 0 {
		if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_DATAA, data); err != nil {
			return err
		}
	}
	if port == 1 {
		if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_DATAB, data); err != nil {
			return err
		}
	}
	return nil
}

func initSX1509(d *SX1509) error {
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_DIRA, 0x00); err != nil {
		return err
	}
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_DIRB, 0x00); err != nil {
		return err
	}
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_OPENDRAINA, 0xFF); err != nil {
		return err
	}
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_OPENDRAINB, 0xFF); err != nil {
		return err
	}
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_POLARITYA, 0xFF); err != nil {
		return err
	}
	if err := d.Dev.WriteByteToReg(d.Address, SX1509_REG_POLARITYB, 0xFF); err != nil {
		return err
	}
	d.initialized = true
	return nil
}
