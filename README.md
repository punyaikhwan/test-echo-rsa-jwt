# test-echo-rsa-jwt

Untuk keperluan testing JWT pada Echo dengan Signing Key RS256.
Lihat [echo-rsa-jwt](https://github.com/punyaikhwan/echo-rsa-jwt)

### Jalankan aplikasi:
```go run main.go```

### Request JWT pada endpoint `/genjwt`:
`curl -X POST -F "name=ikhwan" -F "nim=123456" http://localhost:9000/genjwt`

Respon:
```{"token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0OTM5NzAsIm5hbWUiOiJpa2h3YW4iLCJuaW0iOiIxMjM0NTYifQ.IHeG6YVGN0zZ0rkk65u9pzlP6TCLJYb_9-jUQqZ-FrQLxKjFjpi_kEjONk-iH3ZxJ7r0BkwLBk-0u40-cqkdlSj13pjL-BTwT7ltqJLyndqS2PtvypKLfH2cYGlK63yNhazzhO9OlCWHT8tX70h0FAhv7ZWtVSwvGdm3av95s2CGTZxM8gegfaFMwbkYrFuCGxpuALXEdod_wdoslTFYDoXoJFJ5b2EBTHedj4emUaQVyLv_bEpo_RLhfUkNB7Fgjr0o9O8FpwdPZ4iYxHqEnET5KEMbgzyADeSpo4RGmkT70tfEiyzzH4li59iNVOVMGsoUC-fDkR8RiBXxCjM_zw"}```

### Test JWT pada endpoint `/tesjwt`:
`curl -H "Authorization:Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0OTM5NzAsIm5hbWUiOiJpa2h3YW4iLCJuaW0iOiIxMjM0NTYifQ.IHeG6YVGN0zZ0rkk65u9pzlP6TCLJYb_9-jUQqZ-FrQLxKjFjpi_kEjONk-iH3ZxJ7r0BkwLBk-0u40-cqkdlSj13pjL-BTwT7ltqJLyndqS2PtvypKLfH2cYGlK63yNhazzhO9OlCWHT8tX70h0FAhv7ZWtVSwvGdm3av95s2CGTZxM8gegfaFMwbkYrFuCGxpuALXEdod_wdoslTFYDoXoJFJ5b2EBTHedj4emUaQVyLv_bEpo_RLhfUkNB7Fgjr0o9O8FpwdPZ4iYxHqEnET5KEMbgzyADeSpo4RGmkT70tfEiyzzH4li59iNVOVMGsoUC-fDkR8RiBXxCjM_zw" http://localhost:9000/tesjwt`

Respon:
`{"exp":1619493970,"name":"ikhwan","nim":"123456"}`

Jika JWT tidak valid:
`{"message":"invalid or expired jwt"}`
