/**
 * @Author: xueyanghan
 * @File: random.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/20 14:23
 */

package base

// GenerateRandom return a random bytes with fixed length
func GenerateRandom(c *Ctx, s SessionHandle, length int) ([]byte, error) {
	return c.SDFGenerateRandom(s, uint(length))
}
