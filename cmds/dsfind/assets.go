package main

var (

    // Examples is a map to asset files associated with main package
    Examples = map[string][]byte{
    "index": []byte{0xa,0x23,0x23,0x20,0x45,0x58,0x41,0x4d,0x50,0x4c,0x45,0x53,0xa,0xa,0x49,0x6e,0x20,0x74,0x68,0x65,0x20,0x65,0x78,0x61,0x6d,0x70,0x6c,0x65,0x20,0x74,0x68,0x65,0x20,0x69,0x6e,0x64,0x65,0x78,0x20,0x77,0x69,0x6c,0x6c,0x20,0x62,0x65,0x20,0x63,0x72,0x65,0x61,0x74,0x65,0x64,0x20,0x66,0x6f,0x72,0x20,0x61,0x20,0x63,0x6f,0x6c,0x6c,0x65,0x63,0x74,0x69,0x6f,0x6e,0x20,0x63,0x61,0x6c,0x6c,0x65,0x64,0x20,0x22,0x63,0x68,0x61,0x72,0x61,0x63,0x74,0x65,0x72,0x73,0x22,0x2e,0xa,0xa,0x20,0x20,0x20,0x20,0x64,0x73,0x66,0x69,0x6e,0x64,0x20,0x63,0x68,0x61,0x72,0x61,0x63,0x74,0x65,0x72,0x73,0x2e,0x62,0x6c,0x65,0x76,0x65,0x20,0x22,0x4a,0x61,0x63,0x6b,0x20,0x46,0x6c,0x61,0x6e,0x64,0x65,0x72,0x73,0x22,0xa,0xa,0x54,0x68,0x69,0x73,0x20,0x77,0x6f,0x75,0x6c,0x64,0x20,0x73,0x65,0x61,0x72,0x63,0x68,0x20,0x74,0x68,0x65,0x20,0x42,0x6c,0x65,0x76,0x65,0x20,0x69,0x6e,0x64,0x65,0x78,0x20,0x6e,0x61,0x6d,0x65,0x64,0x20,0x63,0x68,0x61,0x72,0x61,0x63,0x74,0x65,0x72,0x73,0x2e,0x62,0x6c,0x65,0x76,0x65,0x20,0x66,0x6f,0x72,0x20,0x74,0x68,0x65,0x20,0x73,0x74,0x72,0x69,0x6e,0x67,0x20,0x22,0x4a,0x61,0x63,0x6b,0x20,0x46,0x6c,0x61,0x6e,0x64,0x65,0x72,0x73,0x22,0x20,0xa,0x72,0x65,0x74,0x75,0x72,0x6e,0x69,0x6e,0x67,0x20,0x72,0x65,0x63,0x6f,0x72,0x64,0x73,0x20,0x74,0x68,0x61,0x74,0x20,0x6d,0x61,0x74,0x63,0x68,0x65,0x64,0x20,0x62,0x61,0x73,0x65,0x64,0x20,0x6f,0x6e,0x20,0x68,0x6f,0x77,0x20,0x74,0x68,0x65,0x20,0x69,0x6e,0x64,0x65,0x78,0x20,0x77,0x61,0x73,0x20,0x64,0x65,0x66,0x69,0x6e,0x65,0x64,0x2e,0xa,0xa},

    "nav": []byte{0x2b,0x20,0x5b,0x48,0x6f,0x6d,0x65,0x5d,0x28,0x2f,0x29,0xa,0x2b,0x20,0x5b,0x55,0x70,0x5d,0x28,0x2e,0x2e,0x2f,0x29,0xa,0x2b,0x20,0x5b,0x64,0x73,0x66,0x69,0x6e,0x64,0x5d,0x28,0x2e,0x2f,0x29,0xa},

	}
    // Help is a map to asset files associated with main package
    Help = map[string][]byte{
    "description": []byte{0xa,0x23,0x23,0x20,0x44,0x65,0x73,0x63,0x72,0x69,0x70,0x74,0x69,0x6f,0x6e,0xa,0xa,0x64,0x73,0x66,0x69,0x6e,0x64,0x20,0x69,0x73,0x20,0x61,0x20,0x63,0x6f,0x6d,0x6d,0x61,0x6e,0x64,0x20,0x6c,0x69,0x6e,0x65,0x20,0x74,0x6f,0x6f,0x6c,0x20,0x66,0x6f,0x72,0x20,0x71,0x75,0x65,0x72,0x79,0x69,0x6e,0x67,0x20,0x61,0x20,0x42,0x6c,0x65,0x76,0x65,0x20,0x69,0x6e,0x64,0x65,0x78,0x65,0x73,0x20,0x62,0x61,0x73,0x65,0x64,0x20,0x6f,0x6e,0x20,0x74,0x68,0x65,0x20,0x72,0x65,0x63,0x6f,0x72,0x64,0x73,0x20,0x69,0x6e,0x20,0x61,0x20,0xa,0x64,0x61,0x74,0x61,0x73,0x65,0x74,0x20,0x63,0x6f,0x6c,0x6c,0x65,0x63,0x74,0x69,0x6f,0x6e,0x2e,0x20,0x42,0x79,0x20,0x64,0x65,0x66,0x61,0x75,0x6c,0x74,0x20,0x64,0x73,0x66,0x69,0x6e,0x64,0x20,0x69,0x73,0x20,0x61,0x73,0x73,0x75,0x6d,0x65,0x64,0x20,0x74,0x68,0x65,0x72,0x65,0x20,0x69,0x73,0x20,0x61,0x6e,0x20,0x69,0x6e,0x64,0x65,0x78,0x20,0x6e,0x61,0x6d,0x65,0x64,0x20,0x61,0x66,0x74,0x65,0x72,0x20,0x74,0x68,0x65,0x20,0xa,0x63,0x6f,0x6c,0x6c,0x65,0x63,0x74,0x69,0x6f,0x6e,0x2e,0x20,0x41,0x6e,0x20,0x6f,0x70,0x74,0x69,0x6f,0x6e,0x20,0x6c,0x65,0x74,0x73,0x20,0x79,0x6f,0x75,0x20,0x63,0x68,0x6f,0x6f,0x73,0x65,0x20,0x64,0x69,0x66,0x66,0x65,0x72,0x65,0x6e,0x74,0x20,0x69,0x6e,0x64,0x65,0x78,0x65,0x73,0x20,0x74,0x6f,0x20,0x71,0x75,0x65,0x72,0x79,0x2e,0x20,0x52,0x65,0x73,0x75,0x6c,0x74,0x73,0x20,0x61,0x72,0x65,0x20,0xa,0x77,0x72,0x69,0x74,0x74,0x65,0x6e,0x20,0x74,0x6f,0x20,0x73,0x74,0x61,0x6e,0x64,0x61,0x72,0x64,0x20,0x6f,0x75,0x74,0x20,0x61,0x6e,0x64,0x20,0x61,0x72,0x65,0x20,0x70,0x61,0x67,0x65,0x64,0x2e,0x20,0x54,0x68,0x65,0x20,0x71,0x75,0x65,0x72,0x79,0x20,0x73,0x79,0x6e,0x74,0x61,0x78,0x20,0x73,0x75,0x70,0x70,0x6f,0x72,0x74,0x65,0x64,0x20,0x69,0x73,0x20,0x64,0x65,0x73,0x63,0x72,0x69,0x62,0x65,0x64,0xa,0x61,0x74,0x20,0x68,0x74,0x74,0x70,0x3a,0x2f,0x2f,0x77,0x77,0x77,0x2e,0x62,0x6c,0x65,0x76,0x65,0x73,0x65,0x61,0x72,0x63,0x68,0x2e,0x63,0x6f,0x6d,0x2f,0x64,0x6f,0x63,0x73,0x2f,0x51,0x75,0x65,0x72,0x79,0x2d,0x53,0x74,0x72,0x69,0x6e,0x67,0x2d,0x51,0x75,0x65,0x72,0x79,0x2f,0x2e,0xa,0xa,0x4f,0x70,0x74,0x69,0x6f,0x6e,0x73,0x20,0x63,0x61,0x6e,0x20,0x62,0x65,0x20,0x75,0x73,0x65,0x64,0x20,0x74,0x6f,0x20,0x6d,0x6f,0x64,0x69,0x66,0x79,0x20,0x74,0x68,0x65,0x20,0x74,0x79,0x70,0x65,0x20,0x6f,0x66,0x20,0x69,0x6e,0x64,0x65,0x78,0x65,0x73,0x20,0x71,0x75,0x65,0x72,0x69,0x65,0x64,0x20,0x61,0x73,0x20,0x77,0x65,0x6c,0x6c,0x20,0x61,0x73,0x20,0x68,0x6f,0x77,0x20,0x72,0x65,0x73,0x75,0x6c,0x74,0x73,0xa,0x61,0x72,0x65,0x20,0x6f,0x75,0x74,0x70,0x75,0x74,0x2e,0xa,0xa},

    "index": []byte{0xa,0x23,0x20,0x64,0x73,0x66,0x69,0x6e,0x64,0xa,0xa,0x55,0x73,0x65,0x73,0x20,0x42,0x6c,0x65,0x76,0x65,0x20,0x69,0x6e,0x64,0x65,0x78,0x65,0x73,0x20,0x74,0x6f,0x20,0x73,0x65,0x61,0x72,0x63,0x68,0x20,0x5f,0x64,0x61,0x74,0x61,0x73,0x65,0x74,0x5f,0x20,0x63,0x6f,0x6c,0x6c,0x65,0x63,0x74,0x69,0x6f,0x6e,0x73,0x20,0x6f,0x6e,0x20,0x74,0x68,0x65,0x20,0x63,0x6f,0x6d,0x6d,0x61,0x6e,0x64,0x20,0x6c,0x69,0x6e,0x65,0x2e,0xa,0xa,0x2b,0x20,0x5b,0x75,0x73,0x61,0x67,0x65,0x5d,0x28,0x75,0x73,0x61,0x67,0x65,0x2e,0x68,0x74,0x6d,0x6c,0x29,0xa,0x2b,0x20,0x5b,0x64,0x65,0x73,0x63,0x72,0x69,0x70,0x74,0x69,0x6f,0x6e,0x5d,0x28,0x64,0x65,0x73,0x63,0x72,0x69,0x70,0x74,0x69,0x6f,0x6e,0x2e,0x68,0x74,0x6d,0x6c,0x29,0xa,0xa},

    "nav": []byte{0x2b,0x20,0x5b,0x48,0x6f,0x6d,0x65,0x5d,0x28,0x2f,0x29,0xa,0x2b,0x20,0x5b,0x55,0x70,0x5d,0x28,0x2e,0x2e,0x2f,0x29,0xa,0x2b,0x20,0x5b,0x64,0x73,0x66,0x69,0x6e,0x64,0x5d,0x28,0x2e,0x2f,0x29,0xa},

    "usage": []byte{0xa,0x23,0x20,0x55,0x53,0x41,0x47,0x45,0xa,0xa,0x20,0x20,0x20,0x20,0x64,0x73,0x66,0x69,0x6e,0x64,0x20,0x5b,0x4f,0x50,0x54,0x49,0x4f,0x4e,0x53,0x5d,0x20,0x5b,0x49,0x4e,0x44,0x45,0x58,0x5f,0x4c,0x49,0x53,0x54,0x5d,0x20,0x53,0x45,0x41,0x52,0x43,0x48,0x5f,0x53,0x54,0x52,0x49,0x4e,0x47,0x53,0xa,0xa},

	}

)
