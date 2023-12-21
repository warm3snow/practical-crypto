package base

func SM3Hash(c *Ctx, s SessionHandle, origin []byte) ([]byte, error) {
	err := c.SDFHashInit(s, SGD_SM3, nil, 0)
	if err != nil {
		return nil, err
	}

	err = c.SDFHashUpdate(s, origin, uint(len(origin)))
	if err != nil {
		return nil, err
	}

	v, _, err := c.SDFHashFinal(s)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func SM3HMac(c *Ctx, s SessionHandle, keyIndex uint, origin []byte) (mac []byte, err error) {
	k, err := c.SDFGetSymmKeyHandle(s, keyIndex)
	if err != nil {
		return nil, err
	}

	mac, _, err = c.SDFHMAC(s, k, SGD_SM3, origin, uint(len(origin)))
	if err != nil {
		return nil, err
	}
	return mac, nil
}
