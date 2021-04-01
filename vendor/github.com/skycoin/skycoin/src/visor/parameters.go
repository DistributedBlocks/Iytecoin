package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
 */

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 100000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
"PTLEkPECCxp23gjNu6Y7tSJfFSquLSEDwo",
"2cnTbrWMS5N3WfhLHPWZRHzcNf958tE5m7D",
"88FDicKmxuirCepkGHrczkGyRtCXBRZpce",
"2HXq8zZoZazeoPRJM2afSLsgeqGRciKMtus",
"ZznrXvoi4aVFMNxiDssuUaBft2RaEDZVSK",
"Di6r1yo5JxB4WJdvw751bhnpzYKi3jxut5",
"rs7jMkLN3qqZ7kwDwFEiWksoMBrw7evvbv",
"22n8KySta47mwKpcfZHduK1Kbr5aWugf8Yt",
"fvoDbdbAR2AevUatmw3DBWx2EBwiy1p86W",
"sF4LyG1JpNFg7S5LJmbWsw2SkHyB3Q3kkJ",
"2ebQKFv1DWDhbePvSSdj5QuicPV4Jkee4ZQ",
"2ZrUzmS5TFDr9cPQKjUbFTxysSaGCmwcKGV",
"nrnycUgaVThbw1ZisY4VvpNHRmnuQc8Nmh",
"2HQQZ9hek1YFuK1UGREjLg7PhcLzubBfGaV",
"gcSNd9QQQtBPMpQg6nd93TbJ7q4FKf4ckw",
"obQKRQKuwbouDUYgx5hg928xVbz3pKo9kQ",
"2cQ9AT2shrqtZAjipqrf7SwAruDhounrDWi",
"viyUuMK4stdJXGiYVf4LR97K3gZZ1z7CCM",
"gAxULCeHGzMYwL6f29NHUcg6bE6TN2GpS9",
"2U3vPRCLkJS7RBYzMFYeuCMdPsFibPQeD9q",
"FVxFLaW8ytijp74mnDaXFysd8a4huNxu4p",
"2fRMw2wVPPyqG4FJVtnPBmhRzXVzmN17EYt",
"RbyywAGGcuaK1Vb3SZv81Z6HZTxby2P68u",
"RT8nY2jPAWR1tnjJwPR3wjxMkrJbgRAU9Z",
"26xLYLWJUv1C2wqYG7uDSL8rS6RZk4PLpnH",
"JsotGwkhM7ZBsiiiVg3CZ9dwD5pT9YN2FS",
"w2RXS8LJ3WD995giUkuarXP6H4FEUnDzh9",
"ckp8yWDTSNrU3N9isxKWmYDR8sbbBGaL9V",
"kYJtMGBWa8y8Jf4gcnoghRZ2kxj9xTRHjL",
"2YTx2AkszWP4M6df2TWQLJ9K9pD6aH2Vi8g",
"DfTwwJcitYn5nXTdYnxMUigM3kVPjmrwTi",
"22odQDZ3Y1eWzVD7Vriva1Sji9RBnvRGzfH",
"ChwweaXVKiBgb92Ykf5NiGt2KyE1R4rgyq",
"2RcNqVdSmnoqfSPxrUxp3qw4zA1Fea27ejx",
"cLmCUoKJtXG5e2fFzbsPDL8XVDxgbf6AQp",
"3N6PxTXGcnMRsMsv83n9144KQdkM5tNKFx",
"2YKcu97oXBMCMGJAjD71GRopwgfFCaj2svt",
"WjtAPjAdtnQS8oCMuAA5627uLSwrahscGy",
"26UGmnB4ULoH51Kf2ANMkgEGMFV9M5nm3fV",
"2jniyr6TRSnSDhVXb1qaTcGMw1xMKWXj1vV",
"mcx18xLqFh5jGUuB443AeBGaCBEZ8owCz5",
"rEb5kUwH9G9MD71Qe3d4eUHWhdF8s7EMMv",
"gK3cewzDEF8bQdnVHPCs2MSGLibngQGYzv",
"2WZ9WAvx143DBmzuz7p974Tt2m1wktejXZD",
"jJyYwvFPbhy8SwtA1Y3DDuwSACBUfhmeWE",
"94bxAwZRwFRNA22qZ3HBUkHTq9dJAF28bo",
"jQn6obko21mPn42AALpyyEdC7yiiduzfqC",
"xt1gJXgDvbynR9pqDk6zybmmZwRkLckFA6",
"GESiHwDhTx7uEBWXpp1HnxVJe7uzdnZoDz",
"gEkGAWogWCorrex9YF8Nh9kYUD7eBADnc2",
"25H8ZcEHZftZxcuKBpEboUCiEaQBuMkPKmP",
"eP4kPLHwmVzJRi7buPCUSq3iKQVP1jVLe1",
"3HYY19NdqX2X8EVmfGdVdf1pkd9hKP5gWf",
"GFR7PduvzXvzCXR2onkTEteU7LTX8qvfG8",
"2L6RcqmcdhV1T4xCAty51Kp7LP5RS1upzAc",
"2A1ujonxVdErNqvk8rUJFFDZrip7kurBswp",
"3ywkuZ9eFnuR6YFeCk5um454f2hsUGDA1a",
"2YYLNbDAet2QJi275TnjGbbCdVP6EMdXgy2",
"rqYvh2xYybGoS2fn5KuxbzGfBHLzVncgMe",
"2555anRodaW91RZ34XLHbFa8FmmGPmZKc1t",
"LqF6jVvf5GhAmpUKmhhjgGFBHWmYs8ND4S",
"ZVadbG8bUgPEtAhb2yZSjRQmxRVGfZGwcy",
"noDmpiHsKUYMsfBKk56fMVRWKLw2cDfU5d",
"U85oNhVtj6s5ZfSZUHiWSvmZwGmZ2f39a9",
"2GHMyvqm3dpfCr4dhrXS2QuQLsbf1qFR3Lu",
"28x9Fv2vMWEAuHm3wpfQCBDg3FFpzdUkv2W",
"2BZkJeEwWJLQuow7diZRLvpWGXvaF7dGLRR",
"WNbpaT2yHWYR8j1anxBAwkm6Zwep5fKbGj",
"2LsEujdz1a7s3j1RyyT7SXDypZJVMwoRbn4",
"2fR1NZoTEHhYtYSqTvYgQUuH1S8Vn6yoY94",
"MPPLAdhMZaDoEpL1b7rGbJHssNMFpynVaU",
"ypiP3GsBHsdQHdUo3Nzf4f4Tx2fVhjrpdr",
"2e16RuC959qt6kkk9fe7GdR3jL9rrnCZjNs",
"2MYmBQha84yu1pytEA1FhrUKNf4yos2Wnpa",
"8jRcoziTEdfXj7nZEHjiwCs8kK9SrTyPCx",
"24XfS3rm5rSiMmx3XjkCm1fVKttu9dSuhi9",
"2jd8U6AosNM9naScqd2kbEscv2iVfXgevjd",
"z73x1cVW9wqHBy9eo1KnNqcmHa2qFo8sMj",
"2cpgq4kpAfkwcemgLqcXZdMEtUAMdjjhWX1",
"T6qWTsXghb8EdWR7dyjzMvZh3fStdrtEtB",
"yCiyZVQNA6Xrr5FKK224s9VSpYCYFsjezg",
"2Yhik4pEuuez6Ti7KM8oZrtMmdB9BVFT4FG",
"MyrYPJR2SrUtoBBBV2x1hnF5dTWeoqrqoX",
"2Qic5iwCDvtQAyGjLVHKZfMB5hSPhBAAFSt",
"2Kcj1dCLgy4CPi8LxpBHGUK8jP2KhVXiWF6",
"2hKDMT8UjnhiW9VtEPU7SiuTm7ktCPbT7Ge",
"cpYL2ANdkk7dgcU81WawgQWsDQyLyZsJAd",
"2KmDmQwrnKLEMYsZLoYpChvnRkCWdhKFGsF",
"Hyhgc89aSW7vgdv9kjYikJbinK4fRLzk5R",
"zNizRyH89QeMEeYUHcbqbGginSgaiJJgWi",
"UpGbQULEqsFF19BLQrdpovk1PhLFGMPdbx",
"rRSzAVs5q8Yc2u6vVSaujkt2yVGF9DgkSL",
"2anwS8bDiSne8Ctn2kGmKTbJvVZDKkyjdHZ",
"2eZfaFWBobTvVodnQhKLaMZRnBGZsqGJyMZ",
"2SwHaKzGMPuXiEUcj2xUT2Rez1WNUoB2roC",
"EMZXQzDUw1dtc34acJrxqT81mEqKjwtNzF",
"2i9uAM8z3fwyEZycSzf8W6743REcaJoppz",
"2QvAXA2MJcea48kGihXSSUf64194c5o7mTb",
"2G4MPRTLuuPrCuo1UWunxkCc7Pp2Se5D1Qp",
"4wXaJiZkrSd38g8PMQdW7BDWsxn5tVviiv",}
