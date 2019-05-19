// GENERATED FILE. DO NOT EDIT.

package mkvparse

// Supported ElementIDs
const (
	EBMLElement ElementID = 0x1A45DFA3
	EBMLVersionElement ElementID = 0x4286
	EBMLReadVersionElement ElementID = 0x42F7
	EBMLMaxIDLengthElement ElementID = 0x42F2
	EBMLMaxSizeLengthElement ElementID = 0x42F3
	DocTypeElement ElementID = 0x4282
	DocTypeVersionElement ElementID = 0x4287
	DocTypeReadVersionElement ElementID = 0x4285
	VoidElement ElementID = 0xEC
	CRC32Element ElementID = 0xBF
	SignatureSlotElement ElementID = 0x1B538667
	SignatureAlgoElement ElementID = 0x7E8A
	SignatureHashElement ElementID = 0x7E9A
	SignaturePublicKeyElement ElementID = 0x7EA5
	SignatureElement ElementID = 0x7EB5
	SignatureElementsElement ElementID = 0x7E5B
	SignatureElementListElement ElementID = 0x7E7B
	SignedElementElement ElementID = 0x6532
	SegmentElement ElementID = 0x18538067
	SeekHeadElement ElementID = 0x114D9B74
	SeekElement ElementID = 0x4DBB
	SeekIDElement ElementID = 0x53AB
	SeekPositionElement ElementID = 0x53AC
	InfoElement ElementID = 0x1549A966
	SegmentUIDElement ElementID = 0x73A4
	SegmentFilenameElement ElementID = 0x7384
	PrevUIDElement ElementID = 0x3CB923
	PrevFilenameElement ElementID = 0x3C83AB
	NextUIDElement ElementID = 0x3EB923
	NextFilenameElement ElementID = 0x3E83BB
	SegmentFamilyElement ElementID = 0x4444
	ChapterTranslateElement ElementID = 0x6924
	ChapterTranslateEditionUIDElement ElementID = 0x69FC
	ChapterTranslateCodecElement ElementID = 0x69BF
	ChapterTranslateIDElement ElementID = 0x69A5
	TimecodeScaleElement ElementID = 0x2AD7B1
	DurationElement ElementID = 0x4489
	DateUTCElement ElementID = 0x4461
	TitleElement ElementID = 0x7BA9
	MuxingAppElement ElementID = 0x4D80
	WritingAppElement ElementID = 0x5741
	ClusterElement ElementID = 0x1F43B675
	TimecodeElement ElementID = 0xE7
	SilentTracksElement ElementID = 0x5854
	SilentTrackNumberElement ElementID = 0x58D7
	PositionElement ElementID = 0xA7
	PrevSizeElement ElementID = 0xAB
	SimpleBlockElement ElementID = 0xA3
	BlockGroupElement ElementID = 0xA0
	BlockElement ElementID = 0xA1
	BlockVirtualElement ElementID = 0xA2
	BlockAdditionsElement ElementID = 0x75A1
	BlockMoreElement ElementID = 0xA6
	BlockAddIDElement ElementID = 0xEE
	BlockAdditionalElement ElementID = 0xA5
	BlockDurationElement ElementID = 0x9B
	ReferencePriorityElement ElementID = 0xFA
	ReferenceBlockElement ElementID = 0xFB
	ReferenceVirtualElement ElementID = 0xFD
	CodecStateElement ElementID = 0xA4
	DiscardPaddingElement ElementID = 0x75A2
	SlicesElement ElementID = 0x8E
	TimeSliceElement ElementID = 0xE8
	LaceNumberElement ElementID = 0xCC
	FrameNumberElement ElementID = 0xCD
	BlockAdditionIDElement ElementID = 0xCB
	DelayElement ElementID = 0xCE
	SliceDurationElement ElementID = 0xCF
	ReferenceFrameElement ElementID = 0xC8
	ReferenceOffsetElement ElementID = 0xC9
	ReferenceTimeCodeElement ElementID = 0xCA
	EncryptedBlockElement ElementID = 0xAF
	TracksElement ElementID = 0x1654AE6B
	TrackEntryElement ElementID = 0xAE
	TrackNumberElement ElementID = 0xD7
	TrackUIDElement ElementID = 0x73C5
	TrackTypeElement ElementID = 0x83
	FlagEnabledElement ElementID = 0xB9
	FlagDefaultElement ElementID = 0x88
	FlagForcedElement ElementID = 0x55AA
	FlagLacingElement ElementID = 0x9C
	MinCacheElement ElementID = 0x6DE7
	MaxCacheElement ElementID = 0x6DF8
	DefaultDurationElement ElementID = 0x23E383
	DefaultDecodedFieldDurationElement ElementID = 0x234E7A
	TrackTimecodeScaleElement ElementID = 0x23314F
	TrackOffsetElement ElementID = 0x537F
	MaxBlockAdditionIDElement ElementID = 0x55EE
	NameElement ElementID = 0x536E
	LanguageElement ElementID = 0x22B59C
	LanguageIETFElement ElementID = 0x22B59D
	CodecIDElement ElementID = 0x86
	CodecPrivateElement ElementID = 0x63A2
	CodecNameElement ElementID = 0x258688
	AttachmentLinkElement ElementID = 0x7446
	CodecSettingsElement ElementID = 0x3A9697
	CodecInfoURLElement ElementID = 0x3B4040
	CodecDownloadURLElement ElementID = 0x26B240
	CodecDecodeAllElement ElementID = 0xAA
	TrackOverlayElement ElementID = 0x6FAB
	CodecDelayElement ElementID = 0x56AA
	SeekPreRollElement ElementID = 0x56BB
	TrackTranslateElement ElementID = 0x6624
	TrackTranslateEditionUIDElement ElementID = 0x66FC
	TrackTranslateCodecElement ElementID = 0x66BF
	TrackTranslateTrackIDElement ElementID = 0x66A5
	VideoElement ElementID = 0xE0
	FlagInterlacedElement ElementID = 0x9A
	FieldOrderElement ElementID = 0x9D
	StereoModeElement ElementID = 0x53B8
	AlphaModeElement ElementID = 0x53C0
	OldStereoModeElement ElementID = 0x53B9
	PixelWidthElement ElementID = 0xB0
	PixelHeightElement ElementID = 0xBA
	PixelCropBottomElement ElementID = 0x54AA
	PixelCropTopElement ElementID = 0x54BB
	PixelCropLeftElement ElementID = 0x54CC
	PixelCropRightElement ElementID = 0x54DD
	DisplayWidthElement ElementID = 0x54B0
	DisplayHeightElement ElementID = 0x54BA
	DisplayUnitElement ElementID = 0x54B2
	AspectRatioTypeElement ElementID = 0x54B3
	ColourSpaceElement ElementID = 0x2EB524
	GammaValueElement ElementID = 0x2FB523
	FrameRateElement ElementID = 0x2383E3
	ColourElement ElementID = 0x55B0
	MatrixCoefficientsElement ElementID = 0x55B1
	BitsPerChannelElement ElementID = 0x55B2
	ChromaSubsamplingHorzElement ElementID = 0x55B3
	ChromaSubsamplingVertElement ElementID = 0x55B4
	CbSubsamplingHorzElement ElementID = 0x55B5
	CbSubsamplingVertElement ElementID = 0x55B6
	ChromaSitingHorzElement ElementID = 0x55B7
	ChromaSitingVertElement ElementID = 0x55B8
	RangeElement ElementID = 0x55B9
	TransferCharacteristicsElement ElementID = 0x55BA
	PrimariesElement ElementID = 0x55BB
	MaxCLLElement ElementID = 0x55BC
	MaxFALLElement ElementID = 0x55BD
	MasteringMetadataElement ElementID = 0x55D0
	PrimaryRChromaticityXElement ElementID = 0x55D1
	PrimaryRChromaticityYElement ElementID = 0x55D2
	PrimaryGChromaticityXElement ElementID = 0x55D3
	PrimaryGChromaticityYElement ElementID = 0x55D4
	PrimaryBChromaticityXElement ElementID = 0x55D5
	PrimaryBChromaticityYElement ElementID = 0x55D6
	WhitePointChromaticityXElement ElementID = 0x55D7
	WhitePointChromaticityYElement ElementID = 0x55D8
	LuminanceMaxElement ElementID = 0x55D9
	LuminanceMinElement ElementID = 0x55DA
	ProjectionElement ElementID = 0x7670
	ProjectionTypeElement ElementID = 0x7671
	ProjectionPrivateElement ElementID = 0x7672
	ProjectionPoseYawElement ElementID = 0x7673
	ProjectionPosePitchElement ElementID = 0x7674
	ProjectionPoseRollElement ElementID = 0x7675
	AudioElement ElementID = 0xE1
	SamplingFrequencyElement ElementID = 0xB5
	OutputSamplingFrequencyElement ElementID = 0x78B5
	ChannelsElement ElementID = 0x9F
	ChannelPositionsElement ElementID = 0x7D7B
	BitDepthElement ElementID = 0x6264
	TrackOperationElement ElementID = 0xE2
	TrackCombinePlanesElement ElementID = 0xE3
	TrackPlaneElement ElementID = 0xE4
	TrackPlaneUIDElement ElementID = 0xE5
	TrackPlaneTypeElement ElementID = 0xE6
	TrackJoinBlocksElement ElementID = 0xE9
	TrackJoinUIDElement ElementID = 0xED
	TrickTrackUIDElement ElementID = 0xC0
	TrickTrackSegmentUIDElement ElementID = 0xC1
	TrickTrackFlagElement ElementID = 0xC6
	TrickMasterTrackUIDElement ElementID = 0xC7
	TrickMasterTrackSegmentUIDElement ElementID = 0xC4
	ContentEncodingsElement ElementID = 0x6D80
	ContentEncodingElement ElementID = 0x6240
	ContentEncodingOrderElement ElementID = 0x5031
	ContentEncodingScopeElement ElementID = 0x5032
	ContentEncodingTypeElement ElementID = 0x5033
	ContentCompressionElement ElementID = 0x5034
	ContentCompAlgoElement ElementID = 0x4254
	ContentCompSettingsElement ElementID = 0x4255
	ContentEncryptionElement ElementID = 0x5035
	ContentEncAlgoElement ElementID = 0x47E1
	ContentEncKeyIDElement ElementID = 0x47E2
	ContentEncAESSettingsElement ElementID = 0x47E7
	AESSettingsCipherModeElement ElementID = 0x47E8
	ContentSignatureElement ElementID = 0x47E3
	ContentSigKeyIDElement ElementID = 0x47E4
	ContentSigAlgoElement ElementID = 0x47E5
	ContentSigHashAlgoElement ElementID = 0x47E6
	CuesElement ElementID = 0x1C53BB6B
	CuePointElement ElementID = 0xBB
	CueTimeElement ElementID = 0xB3
	CueTrackPositionsElement ElementID = 0xB7
	CueTrackElement ElementID = 0xF7
	CueClusterPositionElement ElementID = 0xF1
	CueRelativePositionElement ElementID = 0xF0
	CueDurationElement ElementID = 0xB2
	CueBlockNumberElement ElementID = 0x5378
	CueCodecStateElement ElementID = 0xEA
	CueReferenceElement ElementID = 0xDB
	CueRefTimeElement ElementID = 0x96
	CueRefClusterElement ElementID = 0x97
	CueRefNumberElement ElementID = 0x535F
	CueRefCodecStateElement ElementID = 0xEB
	AttachmentsElement ElementID = 0x1941A469
	AttachedFileElement ElementID = 0x61A7
	FileDescriptionElement ElementID = 0x467E
	FileNameElement ElementID = 0x466E
	FileMimeTypeElement ElementID = 0x4660
	FileDataElement ElementID = 0x465C
	FileUIDElement ElementID = 0x46AE
	FileReferralElement ElementID = 0x4675
	FileUsedStartTimeElement ElementID = 0x4661
	FileUsedEndTimeElement ElementID = 0x4662
	ChaptersElement ElementID = 0x1043A770
	EditionEntryElement ElementID = 0x45B9
	EditionUIDElement ElementID = 0x45BC
	EditionFlagHiddenElement ElementID = 0x45BD
	EditionFlagDefaultElement ElementID = 0x45DB
	EditionFlagOrderedElement ElementID = 0x45DD
	ChapterAtomElement ElementID = 0xB6
	ChapterUIDElement ElementID = 0x73C4
	ChapterStringUIDElement ElementID = 0x5654
	ChapterTimeStartElement ElementID = 0x91
	ChapterTimeEndElement ElementID = 0x92
	ChapterFlagHiddenElement ElementID = 0x98
	ChapterFlagEnabledElement ElementID = 0x4598
	ChapterSegmentUIDElement ElementID = 0x6E67
	ChapterSegmentEditionUIDElement ElementID = 0x6EBC
	ChapterPhysicalEquivElement ElementID = 0x63C3
	ChapterTrackElement ElementID = 0x8F
	ChapterTrackNumberElement ElementID = 0x89
	ChapterDisplayElement ElementID = 0x80
	ChapStringElement ElementID = 0x85
	ChapLanguageElement ElementID = 0x437C
	ChapLanguageIETFElement ElementID = 0x437D
	ChapCountryElement ElementID = 0x437E
	ChapProcessElement ElementID = 0x6944
	ChapProcessCodecIDElement ElementID = 0x6955
	ChapProcessPrivateElement ElementID = 0x450D
	ChapProcessCommandElement ElementID = 0x6911
	ChapProcessTimeElement ElementID = 0x6922
	ChapProcessDataElement ElementID = 0x6933
	TagsElement ElementID = 0x1254C367
	TagElement ElementID = 0x7373
	TargetsElement ElementID = 0x63C0
	TargetTypeValueElement ElementID = 0x68CA
	TargetTypeElement ElementID = 0x63CA
	TagTrackUIDElement ElementID = 0x63C5
	TagEditionUIDElement ElementID = 0x63C9
	TagChapterUIDElement ElementID = 0x63C4
	TagAttachmentUIDElement ElementID = 0x63C6
	SimpleTagElement ElementID = 0x67C8
	TagNameElement ElementID = 0x45A3
	TagLanguageElement ElementID = 0x447A
	TagLanguageIETFElement ElementID = 0x447B
	TagDefaultElement ElementID = 0x4484
	TagStringElement ElementID = 0x4487
	TagBinaryElement ElementID = 0x4485
)

var elementTypes = map[ElementID]elementType {
	EBMLElement:masterType,
	EBMLVersionElement:uintegerType,
	EBMLReadVersionElement:uintegerType,
	EBMLMaxIDLengthElement:uintegerType,
	EBMLMaxSizeLengthElement:uintegerType,
	DocTypeElement:stringType,
	DocTypeVersionElement:uintegerType,
	DocTypeReadVersionElement:uintegerType,
	VoidElement:binaryType,
	CRC32Element:binaryType,
	SignatureSlotElement:masterType,
	SignatureAlgoElement:uintegerType,
	SignatureHashElement:uintegerType,
	SignaturePublicKeyElement:binaryType,
	SignatureElement:binaryType,
	SignatureElementsElement:masterType,
	SignatureElementListElement:masterType,
	SignedElementElement:binaryType,
	SegmentElement:masterType,
	SeekHeadElement:masterType,
	SeekElement:masterType,
	SeekIDElement:binaryType,
	SeekPositionElement:uintegerType,
	InfoElement:masterType,
	SegmentUIDElement:binaryType,
	SegmentFilenameElement:utf8Type,
	PrevUIDElement:binaryType,
	PrevFilenameElement:utf8Type,
	NextUIDElement:binaryType,
	NextFilenameElement:utf8Type,
	SegmentFamilyElement:binaryType,
	ChapterTranslateElement:masterType,
	ChapterTranslateEditionUIDElement:uintegerType,
	ChapterTranslateCodecElement:uintegerType,
	ChapterTranslateIDElement:binaryType,
	TimecodeScaleElement:uintegerType,
	DurationElement:floatType,
	DateUTCElement:dateType,
	TitleElement:utf8Type,
	MuxingAppElement:utf8Type,
	WritingAppElement:utf8Type,
	ClusterElement:masterType,
	TimecodeElement:uintegerType,
	SilentTracksElement:masterType,
	SilentTrackNumberElement:uintegerType,
	PositionElement:uintegerType,
	PrevSizeElement:uintegerType,
	SimpleBlockElement:binaryType,
	BlockGroupElement:masterType,
	BlockElement:binaryType,
	BlockVirtualElement:binaryType,
	BlockAdditionsElement:masterType,
	BlockMoreElement:masterType,
	BlockAddIDElement:uintegerType,
	BlockAdditionalElement:binaryType,
	BlockDurationElement:uintegerType,
	ReferencePriorityElement:uintegerType,
	ReferenceBlockElement:integerType,
	ReferenceVirtualElement:integerType,
	CodecStateElement:binaryType,
	DiscardPaddingElement:integerType,
	SlicesElement:masterType,
	TimeSliceElement:masterType,
	LaceNumberElement:uintegerType,
	FrameNumberElement:uintegerType,
	BlockAdditionIDElement:uintegerType,
	DelayElement:uintegerType,
	SliceDurationElement:uintegerType,
	ReferenceFrameElement:masterType,
	ReferenceOffsetElement:uintegerType,
	ReferenceTimeCodeElement:uintegerType,
	EncryptedBlockElement:binaryType,
	TracksElement:masterType,
	TrackEntryElement:masterType,
	TrackNumberElement:uintegerType,
	TrackUIDElement:uintegerType,
	TrackTypeElement:uintegerType,
	FlagEnabledElement:uintegerType,
	FlagDefaultElement:uintegerType,
	FlagForcedElement:uintegerType,
	FlagLacingElement:uintegerType,
	MinCacheElement:uintegerType,
	MaxCacheElement:uintegerType,
	DefaultDurationElement:uintegerType,
	DefaultDecodedFieldDurationElement:uintegerType,
	TrackTimecodeScaleElement:floatType,
	TrackOffsetElement:integerType,
	MaxBlockAdditionIDElement:uintegerType,
	NameElement:utf8Type,
	LanguageElement:stringType,
	LanguageIETFElement:stringType,
	CodecIDElement:stringType,
	CodecPrivateElement:binaryType,
	CodecNameElement:utf8Type,
	AttachmentLinkElement:uintegerType,
	CodecSettingsElement:utf8Type,
	CodecInfoURLElement:stringType,
	CodecDownloadURLElement:stringType,
	CodecDecodeAllElement:uintegerType,
	TrackOverlayElement:uintegerType,
	CodecDelayElement:uintegerType,
	SeekPreRollElement:uintegerType,
	TrackTranslateElement:masterType,
	TrackTranslateEditionUIDElement:uintegerType,
	TrackTranslateCodecElement:uintegerType,
	TrackTranslateTrackIDElement:binaryType,
	VideoElement:masterType,
	FlagInterlacedElement:uintegerType,
	FieldOrderElement:uintegerType,
	StereoModeElement:uintegerType,
	AlphaModeElement:uintegerType,
	OldStereoModeElement:uintegerType,
	PixelWidthElement:uintegerType,
	PixelHeightElement:uintegerType,
	PixelCropBottomElement:uintegerType,
	PixelCropTopElement:uintegerType,
	PixelCropLeftElement:uintegerType,
	PixelCropRightElement:uintegerType,
	DisplayWidthElement:uintegerType,
	DisplayHeightElement:uintegerType,
	DisplayUnitElement:uintegerType,
	AspectRatioTypeElement:uintegerType,
	ColourSpaceElement:binaryType,
	GammaValueElement:floatType,
	FrameRateElement:floatType,
	ColourElement:masterType,
	MatrixCoefficientsElement:uintegerType,
	BitsPerChannelElement:uintegerType,
	ChromaSubsamplingHorzElement:uintegerType,
	ChromaSubsamplingVertElement:uintegerType,
	CbSubsamplingHorzElement:uintegerType,
	CbSubsamplingVertElement:uintegerType,
	ChromaSitingHorzElement:uintegerType,
	ChromaSitingVertElement:uintegerType,
	RangeElement:uintegerType,
	TransferCharacteristicsElement:uintegerType,
	PrimariesElement:uintegerType,
	MaxCLLElement:uintegerType,
	MaxFALLElement:uintegerType,
	MasteringMetadataElement:masterType,
	PrimaryRChromaticityXElement:floatType,
	PrimaryRChromaticityYElement:floatType,
	PrimaryGChromaticityXElement:floatType,
	PrimaryGChromaticityYElement:floatType,
	PrimaryBChromaticityXElement:floatType,
	PrimaryBChromaticityYElement:floatType,
	WhitePointChromaticityXElement:floatType,
	WhitePointChromaticityYElement:floatType,
	LuminanceMaxElement:floatType,
	LuminanceMinElement:floatType,
	ProjectionElement:masterType,
	ProjectionTypeElement:uintegerType,
	ProjectionPrivateElement:binaryType,
	ProjectionPoseYawElement:floatType,
	ProjectionPosePitchElement:floatType,
	ProjectionPoseRollElement:floatType,
	AudioElement:masterType,
	SamplingFrequencyElement:floatType,
	OutputSamplingFrequencyElement:floatType,
	ChannelsElement:uintegerType,
	ChannelPositionsElement:binaryType,
	BitDepthElement:uintegerType,
	TrackOperationElement:masterType,
	TrackCombinePlanesElement:masterType,
	TrackPlaneElement:masterType,
	TrackPlaneUIDElement:uintegerType,
	TrackPlaneTypeElement:uintegerType,
	TrackJoinBlocksElement:masterType,
	TrackJoinUIDElement:uintegerType,
	TrickTrackUIDElement:uintegerType,
	TrickTrackSegmentUIDElement:binaryType,
	TrickTrackFlagElement:uintegerType,
	TrickMasterTrackUIDElement:uintegerType,
	TrickMasterTrackSegmentUIDElement:binaryType,
	ContentEncodingsElement:masterType,
	ContentEncodingElement:masterType,
	ContentEncodingOrderElement:uintegerType,
	ContentEncodingScopeElement:uintegerType,
	ContentEncodingTypeElement:uintegerType,
	ContentCompressionElement:masterType,
	ContentCompAlgoElement:uintegerType,
	ContentCompSettingsElement:binaryType,
	ContentEncryptionElement:masterType,
	ContentEncAlgoElement:uintegerType,
	ContentEncKeyIDElement:binaryType,
	ContentEncAESSettingsElement:masterType,
	AESSettingsCipherModeElement:uintegerType,
	ContentSignatureElement:binaryType,
	ContentSigKeyIDElement:binaryType,
	ContentSigAlgoElement:uintegerType,
	ContentSigHashAlgoElement:uintegerType,
	CuesElement:masterType,
	CuePointElement:masterType,
	CueTimeElement:uintegerType,
	CueTrackPositionsElement:masterType,
	CueTrackElement:uintegerType,
	CueClusterPositionElement:uintegerType,
	CueRelativePositionElement:uintegerType,
	CueDurationElement:uintegerType,
	CueBlockNumberElement:uintegerType,
	CueCodecStateElement:uintegerType,
	CueReferenceElement:masterType,
	CueRefTimeElement:uintegerType,
	CueRefClusterElement:uintegerType,
	CueRefNumberElement:uintegerType,
	CueRefCodecStateElement:uintegerType,
	AttachmentsElement:masterType,
	AttachedFileElement:masterType,
	FileDescriptionElement:utf8Type,
	FileNameElement:utf8Type,
	FileMimeTypeElement:stringType,
	FileDataElement:binaryType,
	FileUIDElement:uintegerType,
	FileReferralElement:binaryType,
	FileUsedStartTimeElement:uintegerType,
	FileUsedEndTimeElement:uintegerType,
	ChaptersElement:masterType,
	EditionEntryElement:masterType,
	EditionUIDElement:uintegerType,
	EditionFlagHiddenElement:uintegerType,
	EditionFlagDefaultElement:uintegerType,
	EditionFlagOrderedElement:uintegerType,
	ChapterAtomElement:masterType,
	ChapterUIDElement:uintegerType,
	ChapterStringUIDElement:utf8Type,
	ChapterTimeStartElement:uintegerType,
	ChapterTimeEndElement:uintegerType,
	ChapterFlagHiddenElement:uintegerType,
	ChapterFlagEnabledElement:uintegerType,
	ChapterSegmentUIDElement:binaryType,
	ChapterSegmentEditionUIDElement:uintegerType,
	ChapterPhysicalEquivElement:uintegerType,
	ChapterTrackElement:masterType,
	ChapterTrackNumberElement:uintegerType,
	ChapterDisplayElement:masterType,
	ChapStringElement:utf8Type,
	ChapLanguageElement:stringType,
	ChapLanguageIETFElement:stringType,
	ChapCountryElement:stringType,
	ChapProcessElement:masterType,
	ChapProcessCodecIDElement:uintegerType,
	ChapProcessPrivateElement:binaryType,
	ChapProcessCommandElement:masterType,
	ChapProcessTimeElement:uintegerType,
	ChapProcessDataElement:binaryType,
	TagsElement:masterType,
	TagElement:masterType,
	TargetsElement:masterType,
	TargetTypeValueElement:uintegerType,
	TargetTypeElement:stringType,
	TagTrackUIDElement:uintegerType,
	TagEditionUIDElement:uintegerType,
	TagChapterUIDElement:uintegerType,
	TagAttachmentUIDElement:uintegerType,
	SimpleTagElement:masterType,
	TagNameElement:utf8Type,
	TagLanguageElement:stringType,
	TagLanguageIETFElement:stringType,
	TagDefaultElement:uintegerType,
	TagStringElement:utf8Type,
	TagBinaryElement:binaryType,}

var elementNames = map[ElementID]string {
	EBMLElement: "EBML",
	EBMLVersionElement: "EBMLVersion",
	EBMLReadVersionElement: "EBMLReadVersion",
	EBMLMaxIDLengthElement: "EBMLMaxIDLength",
	EBMLMaxSizeLengthElement: "EBMLMaxSizeLength",
	DocTypeElement: "DocType",
	DocTypeVersionElement: "DocTypeVersion",
	DocTypeReadVersionElement: "DocTypeReadVersion",
	VoidElement: "Void",
	CRC32Element: "CRC32",
	SignatureSlotElement: "SignatureSlot",
	SignatureAlgoElement: "SignatureAlgo",
	SignatureHashElement: "SignatureHash",
	SignaturePublicKeyElement: "SignaturePublicKey",
	SignatureElement: "Signature",
	SignatureElementsElement: "SignatureElements",
	SignatureElementListElement: "SignatureElementList",
	SignedElementElement: "SignedElement",
	SegmentElement: "Segment",
	SeekHeadElement: "SeekHead",
	SeekElement: "Seek",
	SeekIDElement: "SeekID",
	SeekPositionElement: "SeekPosition",
	InfoElement: "Info",
	SegmentUIDElement: "SegmentUID",
	SegmentFilenameElement: "SegmentFilename",
	PrevUIDElement: "PrevUID",
	PrevFilenameElement: "PrevFilename",
	NextUIDElement: "NextUID",
	NextFilenameElement: "NextFilename",
	SegmentFamilyElement: "SegmentFamily",
	ChapterTranslateElement: "ChapterTranslate",
	ChapterTranslateEditionUIDElement: "ChapterTranslateEditionUID",
	ChapterTranslateCodecElement: "ChapterTranslateCodec",
	ChapterTranslateIDElement: "ChapterTranslateID",
	TimecodeScaleElement: "TimecodeScale",
	DurationElement: "Duration",
	DateUTCElement: "DateUTC",
	TitleElement: "Title",
	MuxingAppElement: "MuxingApp",
	WritingAppElement: "WritingApp",
	ClusterElement: "Cluster",
	TimecodeElement: "Timecode",
	SilentTracksElement: "SilentTracks",
	SilentTrackNumberElement: "SilentTrackNumber",
	PositionElement: "Position",
	PrevSizeElement: "PrevSize",
	SimpleBlockElement: "SimpleBlock",
	BlockGroupElement: "BlockGroup",
	BlockElement: "Block",
	BlockVirtualElement: "BlockVirtual",
	BlockAdditionsElement: "BlockAdditions",
	BlockMoreElement: "BlockMore",
	BlockAddIDElement: "BlockAddID",
	BlockAdditionalElement: "BlockAdditional",
	BlockDurationElement: "BlockDuration",
	ReferencePriorityElement: "ReferencePriority",
	ReferenceBlockElement: "ReferenceBlock",
	ReferenceVirtualElement: "ReferenceVirtual",
	CodecStateElement: "CodecState",
	DiscardPaddingElement: "DiscardPadding",
	SlicesElement: "Slices",
	TimeSliceElement: "TimeSlice",
	LaceNumberElement: "LaceNumber",
	FrameNumberElement: "FrameNumber",
	BlockAdditionIDElement: "BlockAdditionID",
	DelayElement: "Delay",
	SliceDurationElement: "SliceDuration",
	ReferenceFrameElement: "ReferenceFrame",
	ReferenceOffsetElement: "ReferenceOffset",
	ReferenceTimeCodeElement: "ReferenceTimeCode",
	EncryptedBlockElement: "EncryptedBlock",
	TracksElement: "Tracks",
	TrackEntryElement: "TrackEntry",
	TrackNumberElement: "TrackNumber",
	TrackUIDElement: "TrackUID",
	TrackTypeElement: "TrackType",
	FlagEnabledElement: "FlagEnabled",
	FlagDefaultElement: "FlagDefault",
	FlagForcedElement: "FlagForced",
	FlagLacingElement: "FlagLacing",
	MinCacheElement: "MinCache",
	MaxCacheElement: "MaxCache",
	DefaultDurationElement: "DefaultDuration",
	DefaultDecodedFieldDurationElement: "DefaultDecodedFieldDuration",
	TrackTimecodeScaleElement: "TrackTimecodeScale",
	TrackOffsetElement: "TrackOffset",
	MaxBlockAdditionIDElement: "MaxBlockAdditionID",
	NameElement: "Name",
	LanguageElement: "Language",
	LanguageIETFElement: "LanguageIETF",
	CodecIDElement: "CodecID",
	CodecPrivateElement: "CodecPrivate",
	CodecNameElement: "CodecName",
	AttachmentLinkElement: "AttachmentLink",
	CodecSettingsElement: "CodecSettings",
	CodecInfoURLElement: "CodecInfoURL",
	CodecDownloadURLElement: "CodecDownloadURL",
	CodecDecodeAllElement: "CodecDecodeAll",
	TrackOverlayElement: "TrackOverlay",
	CodecDelayElement: "CodecDelay",
	SeekPreRollElement: "SeekPreRoll",
	TrackTranslateElement: "TrackTranslate",
	TrackTranslateEditionUIDElement: "TrackTranslateEditionUID",
	TrackTranslateCodecElement: "TrackTranslateCodec",
	TrackTranslateTrackIDElement: "TrackTranslateTrackID",
	VideoElement: "Video",
	FlagInterlacedElement: "FlagInterlaced",
	FieldOrderElement: "FieldOrder",
	StereoModeElement: "StereoMode",
	AlphaModeElement: "AlphaMode",
	OldStereoModeElement: "OldStereoMode",
	PixelWidthElement: "PixelWidth",
	PixelHeightElement: "PixelHeight",
	PixelCropBottomElement: "PixelCropBottom",
	PixelCropTopElement: "PixelCropTop",
	PixelCropLeftElement: "PixelCropLeft",
	PixelCropRightElement: "PixelCropRight",
	DisplayWidthElement: "DisplayWidth",
	DisplayHeightElement: "DisplayHeight",
	DisplayUnitElement: "DisplayUnit",
	AspectRatioTypeElement: "AspectRatioType",
	ColourSpaceElement: "ColourSpace",
	GammaValueElement: "GammaValue",
	FrameRateElement: "FrameRate",
	ColourElement: "Colour",
	MatrixCoefficientsElement: "MatrixCoefficients",
	BitsPerChannelElement: "BitsPerChannel",
	ChromaSubsamplingHorzElement: "ChromaSubsamplingHorz",
	ChromaSubsamplingVertElement: "ChromaSubsamplingVert",
	CbSubsamplingHorzElement: "CbSubsamplingHorz",
	CbSubsamplingVertElement: "CbSubsamplingVert",
	ChromaSitingHorzElement: "ChromaSitingHorz",
	ChromaSitingVertElement: "ChromaSitingVert",
	RangeElement: "Range",
	TransferCharacteristicsElement: "TransferCharacteristics",
	PrimariesElement: "Primaries",
	MaxCLLElement: "MaxCLL",
	MaxFALLElement: "MaxFALL",
	MasteringMetadataElement: "MasteringMetadata",
	PrimaryRChromaticityXElement: "PrimaryRChromaticityX",
	PrimaryRChromaticityYElement: "PrimaryRChromaticityY",
	PrimaryGChromaticityXElement: "PrimaryGChromaticityX",
	PrimaryGChromaticityYElement: "PrimaryGChromaticityY",
	PrimaryBChromaticityXElement: "PrimaryBChromaticityX",
	PrimaryBChromaticityYElement: "PrimaryBChromaticityY",
	WhitePointChromaticityXElement: "WhitePointChromaticityX",
	WhitePointChromaticityYElement: "WhitePointChromaticityY",
	LuminanceMaxElement: "LuminanceMax",
	LuminanceMinElement: "LuminanceMin",
	ProjectionElement: "Projection",
	ProjectionTypeElement: "ProjectionType",
	ProjectionPrivateElement: "ProjectionPrivate",
	ProjectionPoseYawElement: "ProjectionPoseYaw",
	ProjectionPosePitchElement: "ProjectionPosePitch",
	ProjectionPoseRollElement: "ProjectionPoseRoll",
	AudioElement: "Audio",
	SamplingFrequencyElement: "SamplingFrequency",
	OutputSamplingFrequencyElement: "OutputSamplingFrequency",
	ChannelsElement: "Channels",
	ChannelPositionsElement: "ChannelPositions",
	BitDepthElement: "BitDepth",
	TrackOperationElement: "TrackOperation",
	TrackCombinePlanesElement: "TrackCombinePlanes",
	TrackPlaneElement: "TrackPlane",
	TrackPlaneUIDElement: "TrackPlaneUID",
	TrackPlaneTypeElement: "TrackPlaneType",
	TrackJoinBlocksElement: "TrackJoinBlocks",
	TrackJoinUIDElement: "TrackJoinUID",
	TrickTrackUIDElement: "TrickTrackUID",
	TrickTrackSegmentUIDElement: "TrickTrackSegmentUID",
	TrickTrackFlagElement: "TrickTrackFlag",
	TrickMasterTrackUIDElement: "TrickMasterTrackUID",
	TrickMasterTrackSegmentUIDElement: "TrickMasterTrackSegmentUID",
	ContentEncodingsElement: "ContentEncodings",
	ContentEncodingElement: "ContentEncoding",
	ContentEncodingOrderElement: "ContentEncodingOrder",
	ContentEncodingScopeElement: "ContentEncodingScope",
	ContentEncodingTypeElement: "ContentEncodingType",
	ContentCompressionElement: "ContentCompression",
	ContentCompAlgoElement: "ContentCompAlgo",
	ContentCompSettingsElement: "ContentCompSettings",
	ContentEncryptionElement: "ContentEncryption",
	ContentEncAlgoElement: "ContentEncAlgo",
	ContentEncKeyIDElement: "ContentEncKeyID",
	ContentEncAESSettingsElement: "ContentEncAESSettings",
	AESSettingsCipherModeElement: "AESSettingsCipherMode",
	ContentSignatureElement: "ContentSignature",
	ContentSigKeyIDElement: "ContentSigKeyID",
	ContentSigAlgoElement: "ContentSigAlgo",
	ContentSigHashAlgoElement: "ContentSigHashAlgo",
	CuesElement: "Cues",
	CuePointElement: "CuePoint",
	CueTimeElement: "CueTime",
	CueTrackPositionsElement: "CueTrackPositions",
	CueTrackElement: "CueTrack",
	CueClusterPositionElement: "CueClusterPosition",
	CueRelativePositionElement: "CueRelativePosition",
	CueDurationElement: "CueDuration",
	CueBlockNumberElement: "CueBlockNumber",
	CueCodecStateElement: "CueCodecState",
	CueReferenceElement: "CueReference",
	CueRefTimeElement: "CueRefTime",
	CueRefClusterElement: "CueRefCluster",
	CueRefNumberElement: "CueRefNumber",
	CueRefCodecStateElement: "CueRefCodecState",
	AttachmentsElement: "Attachments",
	AttachedFileElement: "AttachedFile",
	FileDescriptionElement: "FileDescription",
	FileNameElement: "FileName",
	FileMimeTypeElement: "FileMimeType",
	FileDataElement: "FileData",
	FileUIDElement: "FileUID",
	FileReferralElement: "FileReferral",
	FileUsedStartTimeElement: "FileUsedStartTime",
	FileUsedEndTimeElement: "FileUsedEndTime",
	ChaptersElement: "Chapters",
	EditionEntryElement: "EditionEntry",
	EditionUIDElement: "EditionUID",
	EditionFlagHiddenElement: "EditionFlagHidden",
	EditionFlagDefaultElement: "EditionFlagDefault",
	EditionFlagOrderedElement: "EditionFlagOrdered",
	ChapterAtomElement: "ChapterAtom",
	ChapterUIDElement: "ChapterUID",
	ChapterStringUIDElement: "ChapterStringUID",
	ChapterTimeStartElement: "ChapterTimeStart",
	ChapterTimeEndElement: "ChapterTimeEnd",
	ChapterFlagHiddenElement: "ChapterFlagHidden",
	ChapterFlagEnabledElement: "ChapterFlagEnabled",
	ChapterSegmentUIDElement: "ChapterSegmentUID",
	ChapterSegmentEditionUIDElement: "ChapterSegmentEditionUID",
	ChapterPhysicalEquivElement: "ChapterPhysicalEquiv",
	ChapterTrackElement: "ChapterTrack",
	ChapterTrackNumberElement: "ChapterTrackNumber",
	ChapterDisplayElement: "ChapterDisplay",
	ChapStringElement: "ChapString",
	ChapLanguageElement: "ChapLanguage",
	ChapLanguageIETFElement: "ChapLanguageIETF",
	ChapCountryElement: "ChapCountry",
	ChapProcessElement: "ChapProcess",
	ChapProcessCodecIDElement: "ChapProcessCodecID",
	ChapProcessPrivateElement: "ChapProcessPrivate",
	ChapProcessCommandElement: "ChapProcessCommand",
	ChapProcessTimeElement: "ChapProcessTime",
	ChapProcessDataElement: "ChapProcessData",
	TagsElement: "Tags",
	TagElement: "Tag",
	TargetsElement: "Targets",
	TargetTypeValueElement: "TargetTypeValue",
	TargetTypeElement: "TargetType",
	TagTrackUIDElement: "TagTrackUID",
	TagEditionUIDElement: "TagEditionUID",
	TagChapterUIDElement: "TagChapterUID",
	TagAttachmentUIDElement: "TagAttachmentUID",
	SimpleTagElement: "SimpleTag",
	TagNameElement: "TagName",
	TagLanguageElement: "TagLanguage",
	TagLanguageIETFElement: "TagLanguageIETF",
	TagDefaultElement: "TagDefault",
	TagStringElement: "TagString",
	TagBinaryElement: "TagBinary",}

