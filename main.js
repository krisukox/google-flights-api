serialize = function (a, b) {
    // void 0 === b && (b = 0);
    // nna();
    // b = kna[b];
    const c = Array(Math.floor(a.length / 3))
        , d = b[64] || "";
    let e = 0
        , f = 0;
    for (; e < a.length - 2; e += 3) {
        var h = a[e]
            , k = a[e + 1]
            , l = a[e + 2]
            , m = b[h >> 2];
        h = b[(h & 3) << 4 | k >> 4];
        k = b[(k & 15) << 2 | l >> 6];
        l = b[l & 63];
        c[f++] = m + h + k + l
    }
    m = 0;
    l = d;
    switch (a.length - e) {
        case 2:
            m = a[e + 1],
                l = b[(m & 15) << 2] || d;
        case 1:
            a = a[e],
                c[f] = b[a >> 2] + b[(a & 3) << 4 | m >> 4] + l + d
    }
    return c.join("")
}
// Wrocław - Praga

// 2023-07-03 2023-07-07
a = [8, 28, 16, 2, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 51, 106, 12, 8, 2, 18, 8, 47, 109, 47, 48, 56, 52, 53, 98, 114, 12, 8, 3, 18, 8, 47, 109,
    47, 48, 53, 121, 119, 103, 26, 40, 18, 10, 50, 48, 50, 51, 45, 48, 55, 45, 48, 55, 106, 12, 8, 3, 18, 8, 47, 109, 47, 48, 53, 121, 119, 103, 114, 12, 8, 2, 18, 8,
    47, 109, 47, 48, 56, 52, 53, 98, 64, 1, 72, 1, 112, 1, 130, 1, 11, 8, 255, 255, 255, 255, 255, 255, 255, 255, 255, 1, 152, 1, 1
];

b = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
    'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_'];

console.log(a.length)
console.log(b.length)

console.log(serialize(a, b));


// /m/0845b

// CBwQAhooEgoyMDIzLTA3LTAzagwIAhIIL20vMDg0NWJyDAgDEggvbS8wNHl3ZxooEgoyMDIzLTA3LTA3agwIAxIIL20vMDR5d2dyDAgCEggvbS8wODQ1YkABSAFwAYIBCwj___________8BmAEB


// CBwQAhooEgoyMDIzLTA3LTAyagwIAhIIL20vMDg0NWJyDAgDEggvbS8wNXl3ZxooEgoyMDIzLTA3LTA3agwIAxIIL20vMDV5d2dyDAgCEggvbS8wODQ1YkABSAFwAYIBCwj___________8BmAEB


