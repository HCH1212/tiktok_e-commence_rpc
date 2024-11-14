// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package payment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *ChargeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ChargeReq[number], err)
}

func (x *ChargeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *ChargeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CardNum, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ChargeReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ChargeReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ChargeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ChargeResp[number], err)
}

func (x *ChargeResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PaymentId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ChargeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ChargeReq) fastWriteField1(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 1, x.GetAmount())
	return offset
}

func (x *ChargeReq) fastWriteField2(buf []byte) (offset int) {
	if x.CardNum == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCardNum())
	return offset
}

func (x *ChargeReq) fastWriteField3(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetOrderId())
	return offset
}

func (x *ChargeReq) fastWriteField4(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 4, x.GetUserId())
	return offset
}

func (x *ChargeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ChargeResp) fastWriteField1(buf []byte) (offset int) {
	if x.PaymentId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetPaymentId())
	return offset
}

func (x *ChargeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ChargeReq) sizeField1() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeFloat(1, x.GetAmount())
	return n
}

func (x *ChargeReq) sizeField2() (n int) {
	if x.CardNum == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCardNum())
	return n
}

func (x *ChargeReq) sizeField3() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetOrderId())
	return n
}

func (x *ChargeReq) sizeField4() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(4, x.GetUserId())
	return n
}

func (x *ChargeResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ChargeResp) sizeField1() (n int) {
	if x.PaymentId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetPaymentId())
	return n
}

var fieldIDToName_ChargeReq = map[int32]string{
	1: "Amount",
	2: "CardNum",
	3: "OrderId",
	4: "UserId",
}

var fieldIDToName_ChargeResp = map[int32]string{
	1: "PaymentId",
}
