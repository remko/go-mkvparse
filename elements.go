// GENERATED FILE. DO NOT EDIT.

package mkvparse

// Supported ElementIDs
const (
	EBMLElement = 0x1A45DFA3
	EBMLVersionElement = 0x4286
	EBMLReadVersionElement = 0x42F7
	EBMLMaxIDLengthElement = 0x42F2
	EBMLMaxSizeLengthElement = 0x42F3
	DocTypeElement = 0x4282
	DocTypeVersionElement = 0x4287
	DocTypeReadVersionElement = 0x4285
	VoidElement = 0xEC
	CRC32Element = 0xBF
	SignatureSlotElement = 0x1B538667
	SignatureAlgoElement = 0x7E8A
	SignatureHashElement = 0x7E9A
	SignaturePublicKeyElement = 0x7EA5
	SignatureElement = 0x7EB5
	SignatureElementsElement = 0x7E5B
	SignatureElementListElement = 0x7E7B
	SignedElementElement = 0x6532
	SegmentElement = 0x18538067
	SeekHeadElement = 0x114D9B74
	SeekElement = 0x4DBB
	SeekIDElement = 0x53AB
	SeekPositionElement = 0x53AC
	InfoElement = 0x1549A966
	SegmentUIDElement = 0x73A4
	SegmentFilenameElement = 0x7384
	PrevUIDElement = 0x3CB923
	PrevFilenameElement = 0x3C83AB
	NextUIDElement = 0x3EB923
	NextFilenameElement = 0x3E83BB
	SegmentFamilyElement = 0x4444
	ChapterTranslateElement = 0x6924
	ChapterTranslateEditionUIDElement = 0x69FC
	ChapterTranslateCodecElement = 0x69BF
	ChapterTranslateIDElement = 0x69A5
	TimecodeScaleElement = 0x2AD7B1
	DurationElement = 0x4489
	DateUTCElement = 0x4461
	TitleElement = 0x7BA9
	MuxingAppElement = 0x4D80
	WritingAppElement = 0x5741
	ClusterElement = 0x1F43B675
	TimecodeElement = 0xE7
	SilentTracksElement = 0x5854
	SilentTrackNumberElement = 0x58D7
	PositionElement = 0xA7
	PrevSizeElement = 0xAB
	SimpleBlockElement = 0xA3
	BlockGroupElement = 0xA0
	BlockElement = 0xA1
	BlockVirtualElement = 0xA2
	BlockAdditionsElement = 0x75A1
	BlockMoreElement = 0xA6
	BlockAddIDElement = 0xEE
	BlockAdditionalElement = 0xA5
	BlockDurationElement = 0x9B
	ReferencePriorityElement = 0xFA
	ReferenceBlockElement = 0xFB
	ReferenceVirtualElement = 0xFD
	CodecStateElement = 0xA4
	DiscardPaddingElement = 0x75A2
	SlicesElement = 0x8E
	TimeSliceElement = 0xE8
	LaceNumberElement = 0xCC
	FrameNumberElement = 0xCD
	BlockAdditionIDElement = 0xCB
	DelayElement = 0xCE
	SliceDurationElement = 0xCF
	ReferenceFrameElement = 0xC8
	ReferenceOffsetElement = 0xC9
	ReferenceTimeCodeElement = 0xCA
	EncryptedBlockElement = 0xAF
	TracksElement = 0x1654AE6B
	TrackEntryElement = 0xAE
	TrackNumberElement = 0xD7
	TrackUIDElement = 0x73C5
	TrackTypeElement = 0x83
	FlagEnabledElement = 0xB9
	FlagDefaultElement = 0x88
	FlagForcedElement = 0x55AA
	FlagLacingElement = 0x9C
	MinCacheElement = 0x6DE7
	MaxCacheElement = 0x6DF8
	DefaultDurationElement = 0x23E383
	DefaultDecodedFieldDurationElement = 0x234E7A
	TrackTimecodeScaleElement = 0x23314F
	TrackOffsetElement = 0x537F
	MaxBlockAdditionIDElement = 0x55EE
	NameElement = 0x536E
	LanguageElement = 0x22B59C
	CodecIDElement = 0x86
	CodecPrivateElement = 0x63A2
	CodecNameElement = 0x258688
	AttachmentLinkElement = 0x7446
	CodecSettingsElement = 0x3A9697
	CodecInfoURLElement = 0x3B4040
	CodecDownloadURLElement = 0x26B240
	CodecDecodeAllElement = 0xAA
	TrackOverlayElement = 0x6FAB
	CodecDelayElement = 0x56AA
	SeekPreRollElement = 0x56BB
	TrackTranslateElement = 0x6624
	TrackTranslateEditionUIDElement = 0x66FC
	TrackTranslateCodecElement = 0x66BF
	TrackTranslateTrackIDElement = 0x66A5
	VideoElement = 0xE0
	FlagInterlacedElement = 0x9A
	FieldOrderElement = 0x9D
	StereoModeElement = 0x53B8
	AlphaModeElement = 0x53C0
	OldStereoModeElement = 0x53B9
	PixelWidthElement = 0xB0
	PixelHeightElement = 0xBA
	PixelCropBottomElement = 0x54AA
	PixelCropTopElement = 0x54BB
	PixelCropLeftElement = 0x54CC
	PixelCropRightElement = 0x54DD
	DisplayWidthElement = 0x54B0
	DisplayHeightElement = 0x54BA
	DisplayUnitElement = 0x54B2
	AspectRatioTypeElement = 0x54B3
	ColourSpaceElement = 0x2EB524
	GammaValueElement = 0x2FB523
	FrameRateElement = 0x2383E3
	ColourElement = 0x55B0
	MatrixCoefficientsElement = 0x55B1
	BitsPerChannelElement = 0x55B2
	ChromaSubsamplingHorzElement = 0x55B3
	ChromaSubsamplingVertElement = 0x55B4
	CbSubsamplingHorzElement = 0x55B5
	CbSubsamplingVertElement = 0x55B6
	ChromaSitingHorzElement = 0x55B7
	ChromaSitingVertElement = 0x55B8
	RangeElement = 0x55B9
	TransferCharacteristicsElement = 0x55BA
	PrimariesElement = 0x55BB
	MaxCLLElement = 0x55BC
	MaxFALLElement = 0x55BD
	MasteringMetadataElement = 0x55D0
	PrimaryRChromaticityXElement = 0x55D1
	PrimaryRChromaticityYElement = 0x55D2
	PrimaryGChromaticityXElement = 0x55D3
	PrimaryGChromaticityYElement = 0x55D4
	PrimaryBChromaticityXElement = 0x55D5
	PrimaryBChromaticityYElement = 0x55D6
	WhitePointChromaticityXElement = 0x55D7
	WhitePointChromaticityYElement = 0x55D8
	LuminanceMaxElement = 0x55D9
	LuminanceMinElement = 0x55DA
	AudioElement = 0xE1
	SamplingFrequencyElement = 0xB5
	OutputSamplingFrequencyElement = 0x78B5
	ChannelsElement = 0x9F
	ChannelPositionsElement = 0x7D7B
	BitDepthElement = 0x6264
	TrackOperationElement = 0xE2
	TrackCombinePlanesElement = 0xE3
	TrackPlaneElement = 0xE4
	TrackPlaneUIDElement = 0xE5
	TrackPlaneTypeElement = 0xE6
	TrackJoinBlocksElement = 0xE9
	TrackJoinUIDElement = 0xED
	TrickTrackUIDElement = 0xC0
	TrickTrackSegmentUIDElement = 0xC1
	TrickTrackFlagElement = 0xC6
	TrickMasterTrackUIDElement = 0xC7
	TrickMasterTrackSegmentUIDElement = 0xC4
	ContentEncodingsElement = 0x6D80
	ContentEncodingElement = 0x6240
	ContentEncodingOrderElement = 0x5031
	ContentEncodingScopeElement = 0x5032
	ContentEncodingTypeElement = 0x5033
	ContentCompressionElement = 0x5034
	ContentCompAlgoElement = 0x4254
	ContentCompSettingsElement = 0x4255
	ContentEncryptionElement = 0x5035
	ContentEncAlgoElement = 0x47E1
	ContentEncKeyIDElement = 0x47E2
	ContentSignatureElement = 0x47E3
	ContentSigKeyIDElement = 0x47E4
	ContentSigAlgoElement = 0x47E5
	ContentSigHashAlgoElement = 0x47E6
	CuesElement = 0x1C53BB6B
	CuePointElement = 0xBB
	CueTimeElement = 0xB3
	CueTrackPositionsElement = 0xB7
	CueTrackElement = 0xF7
	CueClusterPositionElement = 0xF1
	CueRelativePositionElement = 0xF0
	CueDurationElement = 0xB2
	CueBlockNumberElement = 0x5378
	CueCodecStateElement = 0xEA
	CueReferenceElement = 0xDB
	CueRefTimeElement = 0x96
	CueRefClusterElement = 0x97
	CueRefNumberElement = 0x535F
	CueRefCodecStateElement = 0xEB
	AttachmentsElement = 0x1941A469
	AttachedFileElement = 0x61A7
	FileDescriptionElement = 0x467E
	FileNameElement = 0x466E
	FileMimeTypeElement = 0x4660
	FileDataElement = 0x465C
	FileUIDElement = 0x46AE
	FileReferralElement = 0x4675
	FileUsedStartTimeElement = 0x4661
	FileUsedEndTimeElement = 0x4662
	ChaptersElement = 0x1043A770
	EditionEntryElement = 0x45B9
	EditionUIDElement = 0x45BC
	EditionFlagHiddenElement = 0x45BD
	EditionFlagDefaultElement = 0x45DB
	EditionFlagOrderedElement = 0x45DD
	ChapterAtomElement = 0xB6
	ChapterUIDElement = 0x73C4
	ChapterStringUIDElement = 0x5654
	ChapterTimeStartElement = 0x91
	ChapterTimeEndElement = 0x92
	ChapterFlagHiddenElement = 0x98
	ChapterFlagEnabledElement = 0x4598
	ChapterSegmentUIDElement = 0x6E67
	ChapterSegmentEditionUIDElement = 0x6EBC
	ChapterPhysicalEquivElement = 0x63C3
	ChapterTrackElement = 0x8F
	ChapterTrackNumberElement = 0x89
	ChapterDisplayElement = 0x80
	ChapStringElement = 0x85
	ChapLanguageElement = 0x437C
	ChapCountryElement = 0x437E
	ChapProcessElement = 0x6944
	ChapProcessCodecIDElement = 0x6955
	ChapProcessPrivateElement = 0x450D
	ChapProcessCommandElement = 0x6911
	ChapProcessTimeElement = 0x6922
	ChapProcessDataElement = 0x6933
	TagsElement = 0x1254C367
	TagElement = 0x7373
	TargetsElement = 0x63C0
	TargetTypeValueElement = 0x68CA
	TargetTypeElement = 0x63CA
	TagTrackUIDElement = 0x63C5
	TagEditionUIDElement = 0x63C9
	TagChapterUIDElement = 0x63C4
	TagAttachmentUIDElement = 0x63C6
	SimpleTagElement = 0x67C8
	TagNameElement = 0x45A3
	TagLanguageElement = 0x447A
	TagDefaultElement = 0x4484
	TagStringElement = 0x4487
	TagBinaryElement = 0x4485
)

var elementTypes = map[ElementID]ElementType {
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
	TagDefaultElement: "TagDefault",
	TagStringElement: "TagString",
	TagBinaryElement: "TagBinary",}

