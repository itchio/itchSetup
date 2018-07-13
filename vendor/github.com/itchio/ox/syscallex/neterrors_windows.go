// MACHINE GENERATED BY https://github.com/fasterthanlime/neterrors; DO NOT EDIT

// !build windows

package syscallex

import "syscall"

const (
	NERR_NetNotStarted syscall.Errno = 2102
	NERR_UnknownServer syscall.Errno = 2103
	NERR_ShareMem syscall.Errno = 2104
	NERR_NoNetworkResource syscall.Errno = 2105
	NERR_RemoteOnly syscall.Errno = 2106
	NERR_DevNotRedirected syscall.Errno = 2107
	NERR_ServerNotStarted syscall.Errno = 2114
	NERR_ItemNotFound syscall.Errno = 2115
	NERR_UnknownDevDir syscall.Errno = 2116
	NERR_RedirectedPath syscall.Errno = 2117
	NERR_DuplicateShare syscall.Errno = 2118
	NERR_NoRoom syscall.Errno = 2119
	NERR_TooManyItems syscall.Errno = 2121
	NERR_InvalidMaxUsers syscall.Errno = 2122
	NERR_BufTooSmall syscall.Errno = 2123
	NERR_RemoteErr syscall.Errno = 2127
	NERR_LanmanIniError syscall.Errno = 2131
	NERR_NetworkError syscall.Errno = 2136
	NERR_WkstaInconsistentState syscall.Errno = 2137
	NERR_WkstaNotStarted syscall.Errno = 2138
	NERR_BrowserNotStarted syscall.Errno = 2139
	NERR_InternalError syscall.Errno = 2140
	NERR_BadTransactConfig syscall.Errno = 2141
	NERR_InvalidAPI syscall.Errno = 2142
	NERR_BadEventName syscall.Errno = 2143
	NERR_DupNameReboot syscall.Errno = 2144
	NERR_CfgCompNotFound syscall.Errno = 2146
	NERR_CfgParamNotFound syscall.Errno = 2147
	NERR_LineTooLong syscall.Errno = 2149
	NERR_QNotFound syscall.Errno = 2150
	NERR_JobNotFound syscall.Errno = 2151
	NERR_DestNotFound syscall.Errno = 2152
	NERR_DestExists syscall.Errno = 2153
	NERR_QExists syscall.Errno = 2154
	NERR_QNoRoom syscall.Errno = 2155
	NERR_JobNoRoom syscall.Errno = 2156
	NERR_DestNoRoom syscall.Errno = 2157
	NERR_DestIdle syscall.Errno = 2158
	NERR_DestInvalidOp syscall.Errno = 2159
	NERR_ProcNoRespond syscall.Errno = 2160
	NERR_SpoolerNotLoaded syscall.Errno = 2161
	NERR_DestInvalidState syscall.Errno = 2162
	NERR_QinvalidState syscall.Errno = 2163
	NERR_JobInvalidState syscall.Errno = 2164
	NERR_SpoolNoMemory syscall.Errno = 2165
	NERR_DriverNotFound syscall.Errno = 2166
	NERR_DataTypeInvalid syscall.Errno = 2167
	NERR_ProcNotFound syscall.Errno = 2168
	NERR_ServiceTableLocked syscall.Errno = 2180
	NERR_ServiceTableFull syscall.Errno = 2181
	NERR_ServiceInstalled syscall.Errno = 2182
	NERR_ServiceEntryLocked syscall.Errno = 2183
	NERR_ServiceNotInstalled syscall.Errno = 2184
	NERR_BadServiceName syscall.Errno = 2185
	NERR_ServiceCtlTimeout syscall.Errno = 2186
	NERR_ServiceCtlBusy syscall.Errno = 2187
	NERR_BadServiceProgName syscall.Errno = 2188
	NERR_ServiceNotCtrl syscall.Errno = 2189
	NERR_ServiceKillProc syscall.Errno = 2190
	NERR_ServiceCtlNotValid syscall.Errno = 2191
	NERR_NotInDispatchTbl syscall.Errno = 2192
	NERR_BadControlRecv syscall.Errno = 2193
	NERR_ServiceNotStarting syscall.Errno = 2194
	NERR_AlreadyLoggedOn syscall.Errno = 2200
	NERR_NotLoggedOn syscall.Errno = 2201
	NERR_BadUsername syscall.Errno = 2202
	NERR_BadPassword syscall.Errno = 2203
	NERR_UnableToAddName_W syscall.Errno = 2204
	NERR_UnableToAddName_F syscall.Errno = 2205
	NERR_UnableToDelName_W syscall.Errno = 2206
	NERR_UnableToDelName_F syscall.Errno = 2207
	NERR_LogonsPaused syscall.Errno = 2209
	NERR_LogonServerConflict syscall.Errno = 2210
	NERR_LogonNoUserPath syscall.Errno = 2211
	NERR_LogonScriptError syscall.Errno = 2212
	NERR_StandaloneLogon syscall.Errno = 2214
	NERR_LogonServerNotFound syscall.Errno = 2215
	NERR_LogonDomainExists syscall.Errno = 2216
	NERR_NonValidatedLogon syscall.Errno = 2217
	NERR_ACFNotFound syscall.Errno = 2219
	NERR_GroupNotFound syscall.Errno = 2220
	NERR_UserNotFound syscall.Errno = 2221
	NERR_ResourceNotFound syscall.Errno = 2222
	NERR_GroupExists syscall.Errno = 2223
	NERR_UserExists syscall.Errno = 2224
	NERR_ResourceExists syscall.Errno = 2225
	NERR_NotPrimary syscall.Errno = 2226
	NERR_ACFNotLoaded syscall.Errno = 2227
	NERR_ACFNoRoom syscall.Errno = 2228
	NERR_ACFFileIOFail syscall.Errno = 2229
	NERR_ACFTooManyLists syscall.Errno = 2230
	NERR_UserLogon syscall.Errno = 2231
	NERR_ACFNoParent syscall.Errno = 2232
	NERR_CanNotGrowSegment syscall.Errno = 2233
	NERR_SpeGroupOp syscall.Errno = 2234
	NERR_NotInCache syscall.Errno = 2235
	NERR_UserInGroup syscall.Errno = 2236
	NERR_UserNotInGroup syscall.Errno = 2237
	NERR_AccountUndefined syscall.Errno = 2238
	NERR_AccountExpired syscall.Errno = 2239
	NERR_InvalidWorkstation syscall.Errno = 2240
	NERR_InvalidLogonHours syscall.Errno = 2241
	NERR_PasswordExpired syscall.Errno = 2242
	NERR_PasswordCantChange syscall.Errno = 2243
	NERR_PasswordHistConflict syscall.Errno = 2244
	NERR_PasswordTooShort syscall.Errno = 2245
	NERR_PasswordTooRecent syscall.Errno = 2246
	NERR_InvalidDatabase syscall.Errno = 2247
	NERR_DatabaseUpToDate syscall.Errno = 2248
	NERR_SyncRequired syscall.Errno = 2249
	NERR_UseNotFound syscall.Errno = 2250
	NERR_BadAsgType syscall.Errno = 2251
	NERR_DeviceIsShared syscall.Errno = 2252
	NERR_NoComputerName syscall.Errno = 2270
	NERR_MsgAlreadyStarted syscall.Errno = 2271
	NERR_MsgInitFailed syscall.Errno = 2272
	NERR_NameNotFound syscall.Errno = 2273
	NERR_AlreadyForwarded syscall.Errno = 2274
	NERR_AddForwarded syscall.Errno = 2275
	NERR_AlreadyExists syscall.Errno = 2276
	NERR_TooManyNames syscall.Errno = 2277
	NERR_DelComputerName syscall.Errno = 2278
	NERR_LocalForward syscall.Errno = 2279
	NERR_GrpMsgProcessor syscall.Errno = 2280
	NERR_PausedRemote syscall.Errno = 2281
	NERR_BadReceive syscall.Errno = 2282
	NERR_NameInUse syscall.Errno = 2283
	NERR_MsgNotStarted syscall.Errno = 2284
	NERR_NotLocalName syscall.Errno = 2285
	NERR_NoForwardName syscall.Errno = 2286
	NERR_RemoteFull syscall.Errno = 2287
	NERR_NameNotForwarded syscall.Errno = 2288
	NERR_TruncatedBroadcast syscall.Errno = 2289
	NERR_InvalidDevice syscall.Errno = 2294
	NERR_WriteFault syscall.Errno = 2295
	NERR_DuplicateName syscall.Errno = 2297
	NERR_DeleteLater syscall.Errno = 2298
	NERR_IncompleteDel syscall.Errno = 2299
	NERR_MultipleNets syscall.Errno = 2300
	NERR_NetNameNotFound syscall.Errno = 2310
	NERR_DeviceNotShared syscall.Errno = 2311
	NERR_ClientNameNotFound syscall.Errno = 2312
	NERR_FileIdNotFound syscall.Errno = 2314
	NERR_ExecFailure syscall.Errno = 2315
	NERR_TmpFile syscall.Errno = 2316
	NERR_TooMuchData syscall.Errno = 2317
	NERR_DeviceShareConflict syscall.Errno = 2318
	NERR_BrowserTableIncomplete syscall.Errno = 2319
	NERR_NotLocalDomain syscall.Errno = 2320
	NERR_IsDfsShare syscall.Errno = 2321
	NERR_DevInvalidOpCode syscall.Errno = 2331
	NERR_DevNotFound syscall.Errno = 2332
	NERR_DevNotOpen syscall.Errno = 2333
	NERR_BadQueueDevString syscall.Errno = 2334
	NERR_BadQueuePriority syscall.Errno = 2335
	NERR_NoCommDevs syscall.Errno = 2337
	NERR_QueueNotFound syscall.Errno = 2338
	NERR_BadDevString syscall.Errno = 2340
	NERR_BadDev syscall.Errno = 2341
	NERR_InUseBySpooler syscall.Errno = 2342
	NERR_CommDevInUse syscall.Errno = 2343
	NERR_InvalidComputer syscall.Errno = 2351
	NERR_MaxLenExceeded syscall.Errno = 2354
	NERR_BadComponent syscall.Errno = 2356
	NERR_CantType syscall.Errno = 2357
	NERR_TooManyEntries syscall.Errno = 2362
	NERR_ProfileFileTooBig syscall.Errno = 2370
	NERR_ProfileOffset syscall.Errno = 2371
	NERR_ProfileCleanup syscall.Errno = 2372
	NERR_ProfileUnknownCmd syscall.Errno = 2373
	NERR_ProfileLoadErr syscall.Errno = 2374
	NERR_ProfileSaveErr syscall.Errno = 2375
	NERR_LogOverflow syscall.Errno = 2377
	NERR_LogFileChanged syscall.Errno = 2378
	NERR_LogFileCorrupt syscall.Errno = 2379
	NERR_SourceIsDir syscall.Errno = 2380
	NERR_BadSource syscall.Errno = 2381
	NERR_BadDest syscall.Errno = 2382
	NERR_DifferentServers syscall.Errno = 2383
	NERR_RunSrvPaused syscall.Errno = 2385
	NERR_ErrCommRunSrv syscall.Errno = 2389
	NERR_ErrorExecingGhost syscall.Errno = 2391
	NERR_ShareNotFound syscall.Errno = 2392
	NERR_InvalidLana syscall.Errno = 2400
	NERR_OpenFiles syscall.Errno = 2401
	NERR_ActiveConns syscall.Errno = 2402
	NERR_BadPasswordCore syscall.Errno = 2403
	NERR_DevInUse syscall.Errno = 2404
	NERR_LocalDrive syscall.Errno = 2405
	NERR_AlertExists syscall.Errno = 2430
	NERR_TooManyAlerts syscall.Errno = 2431
	NERR_NoSuchAlert syscall.Errno = 2432
	NERR_BadRecipient syscall.Errno = 2433
	NERR_AcctLimitExceeded syscall.Errno = 2434
	NERR_InvalidLogSeek syscall.Errno = 2440
	NERR_BadUasConfig syscall.Errno = 2450
	NERR_InvalidUASOp syscall.Errno = 2451
	NERR_LastAdmin syscall.Errno = 2452
	NERR_DCNotFound syscall.Errno = 2453
	NERR_LogonTrackingError syscall.Errno = 2454
	NERR_NetlogonNotStarted syscall.Errno = 2455
	NERR_CanNotGrowUASFile syscall.Errno = 2456
	NERR_TimeDiffAtDC syscall.Errno = 2457
	NERR_PasswordMismatch syscall.Errno = 2458
	NERR_NoSuchServer syscall.Errno = 2460
	NERR_NoSuchSession syscall.Errno = 2461
	NERR_NoSuchConnection syscall.Errno = 2462
	NERR_TooManyServers syscall.Errno = 2463
	NERR_TooManySessions syscall.Errno = 2464
	NERR_TooManyConnections syscall.Errno = 2465
	NERR_TooManyFiles syscall.Errno = 2466
	NERR_NoAlternateServers syscall.Errno = 2467
	NERR_TryDownLevel syscall.Errno = 2470
	NERR_UPSDriverNotStarted syscall.Errno = 2480
	NERR_UPSInvalidConfig syscall.Errno = 2481
	NERR_UPSInvalidCommPort syscall.Errno = 2482
	NERR_UPSSignalAsserted syscall.Errno = 2483
	NERR_UPSShutdownFailed syscall.Errno = 2484
	NERR_BadDosRetCode syscall.Errno = 2500
	NERR_ProgNeedsExtraMem syscall.Errno = 2501
	NERR_BadDosFunction syscall.Errno = 2502
	NERR_RemoteBootFailed syscall.Errno = 2503
	NERR_BadFileCheckSum syscall.Errno = 2504
	NERR_NoRplBootSystem syscall.Errno = 2505
	NERR_RplLoadrNetBiosErr syscall.Errno = 2506
	NERR_RplLoadrDiskErr syscall.Errno = 2507
	NERR_ImageParamErr syscall.Errno = 2508
	NERR_TooManyImageParams syscall.Errno = 2509
	NERR_NonDosFloppyUsed syscall.Errno = 2510
	NERR_RplBootRestart syscall.Errno = 2511
	NERR_RplSrvrCallFailed syscall.Errno = 2512
	NERR_CantConnectRplSrvr syscall.Errno = 2513
	NERR_CantOpenImageFile syscall.Errno = 2514
	NERR_CallingRplSrvr syscall.Errno = 2515
	NERR_StartingRplBoot syscall.Errno = 2516
	NERR_RplBootServiceTerm syscall.Errno = 2517
	NERR_RplBootStartFailed syscall.Errno = 2518
	NERR_RPL_CONNECTED syscall.Errno = 2519
	NERR_BrowserConfiguredToNotRun syscall.Errno = 2550
	NERR_RplNoAdaptersStarted syscall.Errno = 2610
	NERR_RplBadRegistry syscall.Errno = 2611
	NERR_RplBadDatabase syscall.Errno = 2612
	NERR_RplRplfilesShare syscall.Errno = 2613
	NERR_RplNotRplServer syscall.Errno = 2614
	NERR_RplCannotEnum syscall.Errno = 2615
	NERR_RplWkstaInfoCorrupted syscall.Errno = 2616
	NERR_RplWkstaNotFound syscall.Errno = 2617
	NERR_RplWkstaNameUnavailable syscall.Errno = 2618
	NERR_RplProfileInfoCorrupted syscall.Errno = 2619
	NERR_RplProfileNotFound syscall.Errno = 2620
	NERR_RplProfileNameUnavailable syscall.Errno = 2621
	NERR_RplProfileNotEmpty syscall.Errno = 2622
	NERR_RplConfigInfoCorrupted syscall.Errno = 2623
	NERR_RplConfigNotFound syscall.Errno = 2624
	NERR_RplAdapterInfoCorrupted syscall.Errno = 2625
	NERR_RplInternal syscall.Errno = 2626
	NERR_RplVendorInfoCorrupted syscall.Errno = 2627
	NERR_RplBootInfoCorrupted syscall.Errno = 2628
	NERR_RplWkstaNeedsUserAcct syscall.Errno = 2629
	NERR_RplNeedsRPLUSERAcct syscall.Errno = 2630
	NERR_RplBootNotFound syscall.Errno = 2631
	NERR_RplIncompatibleProfile syscall.Errno = 2632
	NERR_RplAdapterNameUnavailable syscall.Errno = 2633
	NERR_RplConfigNotEmpty syscall.Errno = 2634
	NERR_RplBootInUse syscall.Errno = 2635
	NERR_RplBackupDatabase syscall.Errno = 2636
	NERR_RplAdapterNotFound syscall.Errno = 2637
	NERR_RplVendorNotFound syscall.Errno = 2638
	NERR_RplVendorNameUnavailable syscall.Errno = 2639
	NERR_RplBootNameUnavailable syscall.Errno = 2640
	NERR_RplConfigNameUnavailable syscall.Errno = 2641
	NERR_DfsInternalCorruption syscall.Errno = 2660
	NERR_DfsVolumeDataCorrupt syscall.Errno = 2661
	NERR_DfsNoSuchVolume syscall.Errno = 2662
	NERR_DfsVolumeAlreadyExists syscall.Errno = 2663
	NERR_DfsAlreadyShared syscall.Errno = 2664
	NERR_DfsNoSuchShare syscall.Errno = 2665
	NERR_DfsNotALeafVolume syscall.Errno = 2666
	NERR_DfsLeafVolume syscall.Errno = 2667
	NERR_DfsVolumeHasMultipleServers syscall.Errno = 2668
	NERR_DfsCantCreateJunctionPoint syscall.Errno = 2669
	NERR_DfsServerNotDfsAware syscall.Errno = 2670
	NERR_DfsBadRenamePath syscall.Errno = 2671
	NERR_DfsVolumeIsOffline syscall.Errno = 2672
	NERR_DfsNoSuchServer syscall.Errno = 2673
	NERR_DfsCyclicalName syscall.Errno = 2674
	NERR_DfsNotSupportedInServerDfs syscall.Errno = 2675
	NERR_DfsDuplicateService syscall.Errno = 2676
	NERR_DfsCantRemoveLastServerShare syscall.Errno = 2677
	NERR_DfsVolumeIsInterDfs syscall.Errno = 2678
	NERR_DfsInconsistent syscall.Errno = 2679
	NERR_DfsServerUpgraded syscall.Errno = 2680
	NERR_DfsDataIsIdentical syscall.Errno = 2681
	NERR_DfsCantRemoveDfsRoot syscall.Errno = 2682
	NERR_DfsChildOrParentInDfs syscall.Errno = 2683
	NERR_DfsInternalError syscall.Errno = 2690
	NERR_SetupAlreadyJoined syscall.Errno = 2691
	NERR_SetupNotJoined syscall.Errno = 2692
	NERR_SetupDomainController syscall.Errno = 2693
	NERR_DefaultJoinRequired syscall.Errno = 2694
	NERR_InvalidWorkgroupName syscall.Errno = 2695
	NERR_NameUsesIncompatibleCodePage syscall.Errno = 2696
	NERR_ComputerAccountNotFound syscall.Errno = 2697
	NERR_PersonalSku syscall.Errno = 2698
	NERR_PasswordMustChange syscall.Errno = 2701
	NERR_AccountLockedOut syscall.Errno = 2702
	NERR_PasswordTooLong syscall.Errno = 2703
	NERR_PasswordNotComplexEnough syscall.Errno = 2704
	NERR_PasswordFilterError syscall.Errno = 2705
	NERR_NoOfflineJoinInfo syscall.Errno = 2709
	NERR_BadOfflineJoinInfo syscall.Errno = 2710
	NERR_CantCreateJoinInfo syscall.Errno = 2711
	NERR_BadDomainJoinInfo syscall.Errno = 2712
	NERR_JoinPerformedMustRestart syscall.Errno = 2713
	NERR_NoJoinPending syscall.Errno = 2714
	NERR_ValuesNotSet syscall.Errno = 2715
	NERR_CantVerifyHostname syscall.Errno = 2716
	NERR_CantLoadOfflineHive syscall.Errno = 2717
	NERR_ConnectionInsecure syscall.Errno = 2718
	NERR_ProvisioningBlobUnsupported syscall.Errno = 2719
)

