package kokomi // 导入yuan-shen模块,需要定期手动更新部分

// FindName 角色外号添加
func FindName(a string) string {
	switch a {
	case "公子":
		a = "达达利亚"
	case "绫华", "神里":
		a = "神里绫华"
	case "绫人":
		a = "神里绫人"
	case "万叶":
		a = "枫原万叶"
	case "雷神":
		a = "雷电将军"
	case "一斗":
		a = "荒泷一斗"
	case "八重", "神子":
		a = "八重神子"
	case "九条":
		a = "九条裟罗"
	case "罗莎":
		a = "罗莎莉亚"
	case "平藏":
		a = "鹿野院平藏"
	case "心海":
		a = "珊瑚宫心海"
	case "散兵":
		a = "流浪者"
	case "草神":
		a = "纳西妲"
	case "风神":
		a = "温迪"
	case "岩神", "帝君":
		a = "钟离"
	}
	return a
}

// Wifequanmap 角色圣遗物评分权重图
var Wifequanmap = map[string]wifequan{
	"神里绫华":  {0, 75, 0, 100, 100, 0, 100, 0, 30, 0},
	"旅行者":   {0, 75, 0, 100, 100, 0, 100, 0, 55, 0},
	"香菱":    {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"菲谢尔":   {0, 75, 0, 100, 100, 75, 100, 60, 0, 0},
	"丽莎":    {0, 75, 0, 100, 100, 75, 100, 0, 0, 0},
	"阿贝多":   {0, 0, 100, 100, 100, 0, 100, 0, 0, 0},
	"安柏":    {0, 75, 0, 100, 100, 75, 100, 100, 0, 0},
	"鹿野院平藏": {0, 75, 0, 100, 100, 75, 100, 0, 30, 0},
	"坎蒂丝":   {75, 75, 0, 100, 100, 0, 100, 0, 55, 0},
	"瑶瑶":    {75, 75, 0, 100, 100, 0, 75, 0, 75, 75},
	"珊瑚宫心海": {100, 50, 0, 0, 0, 0, 100, 0, 55, 100},
	"优菈":    {0, 75, 0, 100, 100, 0, 40, 100, 55, 0},
	"砂糖":    {0, 75, 0, 100, 100, 100, 75, 0, 55, 0},
	"柯莱":    {0, 75, 0, 100, 100, 75, 100, 0, 30, 0},
	"重云":    {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"凯亚":    {0, 75, 0, 100, 100, 0, 100, 100, 30, 0},
	"荒泷一斗":  {0, 50, 100, 100, 100, 0, 100, 0, 30, 0},
	"温迪":    {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"达达利亚":  {0, 75, 0, 100, 100, 75, 100, 0, 30, 0},
	"魈":     {0, 75, 0, 100, 100, 0, 100, 0, 55, 0},
	"托马":    {90, 55, 0, 90, 90, 0, 90, 0, 55, 0},
	"五郎":    {0, 75, 100, 50, 50, 0, 100, 0, 75, 0},
	"班尼特":   {100, 50, 0, 100, 100, 0, 80, 0, 55, 100},
	"刻晴":    {0, 75, 0, 100, 100, 75, 100, 100, 0, 0},
	"辛焱":    {0, 75, 75, 100, 100, 0, 100, 100, 0, 0},
	"夜兰":    {80, 0, 0, 100, 100, 0, 100, 0, 55, 0},
	"纳西妲":   {0, 55, 0, 100, 100, 100, 100, 0, 55, 0},
	"行秋":    {0, 75, 0, 100, 100, 0, 100, 0, 75, 0},
	"胡桃":    {80, 50, 0, 100, 100, 75, 100, 0, 0, 0},
	"迪卢克":   {0, 75, 0, 100, 100, 75, 100, 0, 0, 0},
	"七七":    {0, 75, 0, 100, 100, 0, 100, 100, 55, 100},
	"八重神子":  {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"宵宫":    {0, 75, 0, 100, 100, 75, 100, 0, 0, 0},
	"琴":     {0, 75, 0, 100, 100, 0, 100, 100, 55, 100},
	"空":     {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"提纳里":   {0, 75, 0, 100, 100, 90, 100, 0, 30, 0},
	"枫原万叶":  {0, 75, 0, 100, 100, 100, 100, 0, 55, 0},
	"钟离":    {80, 75, 0, 100, 100, 0, 100, 50, 55, 0},
	"莫娜":    {0, 75, 0, 100, 100, 75, 100, 0, 75, 0},
	"北斗":    {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"凝光":    {0, 75, 0, 100, 100, 0, 100, 0, 30, 0},
	"早柚":    {0, 75, 0, 100, 100, 100, 100, 0, 55, 100},
	"赛诺":    {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"神里绫人":  {50, 75, 0, 100, 100, 0, 100, 0, 30, 0},
	"九条裟罗":  {0, 75, 0, 100, 100, 0, 100, 0, 55, 0},
	"可莉":    {0, 75, 0, 100, 100, 75, 100, 0, 30, 0},
	"艾尔海森":  {0, 50, 0, 100, 100, 80, 100, 0, 35, 0},
	"云堇":    {0, 75, 100, 80, 80, 0, 80, 0, 80, 0},
	"迪奥娜":   {100, 50, 0, 50, 50, 0, 100, 0, 90, 100},
	"荧":     {0, 75, 0, 100, 100, 75, 100, 0, 55, 0},
	"妮露":    {100, 0, 0, 100, 100, 80, 100, 0, 30, 0},
	"申鹤":    {0, 100, 0, 100, 100, 0, 100, 0, 55, 0},
	"埃洛伊":   {0, 75, 0, 100, 100, 0, 100, 0, 0, 0},
	"诺艾尔":   {0, 50, 90, 100, 100, 0, 100, 0, 70, 0},
	"久岐忍":   {100, 50, 0, 100, 100, 75, 100, 0, 55, 100},
	"多莉":    {75, 75, 0, 100, 100, 0, 75, 0, 55, 100},
	"莱依拉":   {100, 75, 0, 100, 100, 0, 100, 0, 35, 0},
	"珐露珊":   {0, 75, 0, 100, 100, 0, 75, 0, 75, 0},
	"罗莎莉亚":  {0, 75, 0, 100, 100, 0, 70, 80, 30, 0},
	"烟绯":    {0, 75, 0, 100, 100, 75, 100, 0, 30, 0},
	"雷泽":    {0, 75, 0, 100, 100, 0, 100, 100, 0, 0},
	"流浪者":   {0, 80, 0, 100, 100, 0, 100, 0, 35, 0},
	"雷电将军":  {0, 75, 0, 100, 100, 0, 75, 0, 90, 0},
	"甘雨":    {0, 75, 0, 100, 100, 75, 100, 0, 0, 0},
}

// Namemap wifename->wifeid
var Namemap = map[string]int64{
	"多莉":    10000068,
	"凯亚":    10000015,
	"宵宫":    10000049,
	"夜兰":    10000060,
	"莱依拉":   10000074,
	"艾尔海森":  10000078,
	"重云":    10000036,
	"辛焱":    10000044,
	"班尼特":   10000032,
	"胡桃":    10000046,
	"云堇":    10000064,
	"久岐忍":   10000065,
	"空":     10000005,
	"香菱":    10000023,
	"早柚":    10000053,
	"神里绫人":  10000066,
	"达达利亚":  10000033,
	"刻晴":    10000042,
	"钟离":    10000030,
	"迪奥娜":   10000039,
	"优菈":    10000051,
	"九条裟罗":  10000056,
	"安柏":    10000021,
	"北斗":    10000024,
	"申鹤":    10000063,
	"纳西妲":   10000073,
	"砂糖":    10000043,
	"荒泷一斗":  10000057,
	"坎蒂丝":   10000072,
	"雷电将军":  10000052,
	"鹿野院平藏": 10000059,
	"诺艾尔":   10000034,
	"甘雨":    10000037,
	"莫娜":    10000041,
	"八重神子":  10000058,
	"柯莱":    10000067,
	"妮露":    10000070,
	"神里绫华":  10000002,
	"丽莎":    10000006,
	"流浪者":   10000075,
	"迪卢克":   10000016,
	"烟绯":    10000048,
	"雷泽":    10000020,
	"阿贝多":   10000038,
	"魈":     10000026,
	"芭芭拉":   10000014,
	"菲谢尔":   10000031,
	"托马":    10000050,
	"珊瑚宫心海": 10000054,
	"提纳里":   10000069,
	"赛诺":    10000071,
	"琴":     10000003,
	"荧":     10000007,
	"珐露珊":   10000076,
	"罗莎莉亚":  10000045,
	"五郎":    10000055,
	"埃洛伊":   10000062,
	"瑶瑶":    10000077,
	"温迪":    10000022,
	"凝光":    10000027,
	"可莉":    10000029,
	"七七":    10000035,
	"枫原万叶":  10000047,
	"凯特":    10000001,
	"行秋":    10000025,
}

// Idtotalent id->talentid
var IdtoTalent = map[int64][]int{
	10000057: []int{10571, 10572, 10575},
	10000014: []int{10070, 10071, 10072},
	10000016: []int{10160, 10161, 10165},
	10000027: []int{10271, 10272, 10274},
	10000050: []int{10501, 10502, 10505},
	10000052: []int{10521, 10522, 10525},
	10000003: []int{10031, 10033, 10034},
	10000043: []int{10431, 10432, 10435},
	10000060: []int{10606, 10607, 10610},
	10000062: []int{10621, 10622, 10625},
	10000024: []int{10241, 10242, 10245},
	10000069: []int{10691, 10692, 10695},
	10000073: []int{10731, 10732, 10735},
	10000020: []int{10201, 10202, 10203},
	10000064: []int{10641, 10642, 10643},
	10000071: []int{10711, 10712, 10715},
	10000074: []int{10741, 10742, 10745},
	10000034: []int{10341, 10342, 10343},
	10000046: []int{10461, 10462, 10463},
	10000053: []int{10531, 10532, 10535},
	10000070: []int{10701, 10702, 10705},
	10000076: []int{10761, 10762, 10765},
	10000015: []int{10073, 10074, 10075},
	10000036: []int{10401, 10402, 10403},
	10000038: []int{10386, 10387, 10388},
	10000044: []int{10441, 10442, 10443},
	10000054: []int{10541, 10542, 10545},
	10000059: []int{10591, 10592, 10595},
	10000002: []int{10024, 10018, 10019},
	10000006: []int{10060, 10061, 10062},
	10000021: []int{10041, 10032, 10017},
	10000025: []int{10381, 10382, 10385},
	10000030: []int{10301, 10302, 10303},
	10000063: []int{10631, 10632, 10635},
	10000067: []int{10671, 10672, 10675},
	10000075: []int{10751, 10752, 10755},
	10000065: []int{10651, 10652, 10655},
	10000068: []int{10681, 10682, 10685},
	10000031: []int{10311, 10312, 10313},
	10000037: []int{10371, 10372, 10373},
	10000041: []int{10411, 10412, 10415},
	10000049: []int{10491, 10492, 10495},
	10000058: []int{10581, 10582, 10585},
	10000032: []int{10321, 10322, 10323},
	10000056: []int{10561, 10562, 10565},
	10000005: []int{100543, 10067, 10068},
	10000048: []int{10481, 10482, 10485},
	10000051: []int{10511, 10512, 10515},
	10000007: []int{100553, 10067, 10068},
	10000023: []int{10231, 10232, 10235},
	10000022: []int{10221, 10224, 10225},
	10000042: []int{10421, 10422, 10425},
	10000055: []int{10551, 10552, 10555},
	10000066: []int{10661, 10662, 10665},
	10000072: []int{10721, 10722, 10725},
	10000033: []int{10331, 10332, 10333},
	10000035: []int{10351, 10352, 10353},
	10000039: []int{10391, 10392, 10395},
	10000045: []int{10451, 10452, 10453},
	10000047: []int{10471, 10472, 10475},
	10000026: []int{10261, 10262, 10265},
	10000029: []int{10291, 10292, 10295},
}

// Promap 角色id匹配属性
var Promap = map[int64]string{
	10000036: "冰",
	10000050: "火",
	10000051: "冰",
	10000066: "水",
	10000067: "草",
	10000016: "火",
	10000025: "水",
	10000030: "岩",
	10000053: "风",
	10000071: "雷",
	10000002: "冰",
	10000003: "风",
	10000005: "风", // 空
	10000068: "雷",
	10000070: "水",
	10000072: "水",
	10000001: "风", // 凯特
	10000055: "岩",
	10000060: "水",
	10000023: "火",
	10000042: "雷",
	10000039: "冰",
	10000057: "岩",
	10000069: "草",
	10000075: "风",
	10000078: "草",
	10000006: "雷",
	10000014: "水",
	10000049: "火",
	10000056: "雷",
	10000058: "雷",
	10000065: "雷",
	10000044: "火",
	10000047: "风",
	10000046: "火",
	10000048: "火",
	10000063: "冰",
	10000076: "风",
	10000015: "冰",
	10000020: "雷",
	10000074: "冰",
	10000022: "风",
	10000064: "岩",
	10000034: "岩",
	10000077: "草",
	10000021: "火",
	10000033: "水",
	10000043: "风",
	10000029: "火",
	10000045: "冰",
	10000062: "冰",
	10000035: "冰",
	10000052: "雷",
	10000031: "雷",
	10000037: "冰",
	10000038: "岩",
	10000041: "水",
	10000007: "风",
	10000024: "雷",
	10000054: "水",
	10000026: "风",
	10000027: "岩",
	10000073: "草",
	10000032: "火",
	10000059: "风",
}

// SywtoName sywname->syw_allname
var SywNamemap = map[string][]string{
	"行者之心":    []string{"故人之心", "归乡之羽", "逐光之石", "异国之盏", "感别之冠"},
	"渡过烈火的贤人": []string{"渡火者的决绝", "渡火者的解脱", "渡火者的煎熬", "渡火者的醒悟", "渡火者的智慧"},
	"昔日宗室之仪":  []string{"宗室之花", "宗室之翎", "宗室时计", "宗室银瓮", "宗室面具"},
	"华馆梦醒形骸记": []string{"荣花之期", "华馆之羽", "众生之谣", "梦醒之瓢", "形骸之笠"},
	"守护之心":    []string{"守护之花", "守护徽印", "守护座钟", "守护之皿", "守护束带"},
	"战狂":      []string{"战狂的蔷薇", "战狂的翎羽", "战狂的时计", "战狂的骨杯", "战狂的鬼面"},
	"幸运儿":     []string{"幸运儿绿花", "幸运儿鹰羽", "幸运儿沙漏", "幸运儿之杯", "幸运儿银冠"},
	"平息鸣雷的尊者": []string{"平雷之心", "平雷之羽", "平雷之刻", "平雷之器", "平雷之冠"},
	"奇迹":      []string{"奇迹之花", "奇迹之羽", "奇迹之沙", "奇迹之杯", "奇迹耳坠"},
	"武人":      []string{"武人的红花", "武人的羽饰", "武人的水漏", "武人的酒杯", "武人的头巾"},
	"翠绿之影":    []string{"野花记忆的绿野", "猎人青翠的箭羽", "翠绿猎人的笃定", "翠绿猎人的容器", "翠绿的猎人之冠"},
	"炽烈的炎之魔女": []string{"魔女的炎之花", "魔女常燃之羽", "魔女破灭之时", "魔女的心之火", "焦灼的魔女帽"},
	"悠古的磐岩":   []string{"磐陀裂生之花", "嵯峨群峰之翼", "星罗圭璧之晷", "巉岩琢塑之樽", "不动玄石之相"},
	"逆飞的流星":   []string{"夏祭之花", "夏祭终末", "夏祭之刻", "夏祭水玉", "夏祭之面"},
	"苍白之火":    []string{"无垢之花", "贤医之羽", "停摆之刻", "超越之盏", "嗤笑之面"},
	"来歆余响":    []string{"魂香之花", "垂玉之叶", "祝祀之凭", "涌泉之盏", "浮溯之珏"},
	"沙上楼阁史话":  []string{"众王之都的开端", "黄金邦国的结末", "失落迷途的机芯", "迷醉长梦的守护", "流沙贵嗣的遗宝"},
	"学士":      []string{"学士的书签", "学士的羽笔", "学士的时钟", "学士的墨杯", "学士的镜片"},
	"游医":      []string{"游医的银莲", "游医的枭羽", "游医的怀钟", "游医的药壶", "游医的方巾"},
	"流浪大地的乐团": []string{"乐团的晨光", "琴师的箭羽", "终幕的时计", "吟游者之壶", "指挥的礼帽"},
	"沉沦之心":    []string{"饰金胸花", "追忆之风", "坚铜罗盘", "沉波之盏", "酒渍船帽"},
	"海染砗磲":    []string{"海染之花", "渊宫之羽", "离别之贝", "真珠之笼", "海祇之冠"},
	"教官":      []string{"教官的胸花", "教官的羽饰", "教官的怀表", "教官的茶杯", "教官的帽子"},
	"祭雷之人":    []string{"", "", "", "", "祭雷礼冠"},
	"辰砂往生录":   []string{"生灵之华", "潜光片羽", "阳辔之遗", "结契之刻", "虺雷之姿"},
	"冒险家":     []string{"冒险家之花", "冒险家尾羽", "冒险家怀表", "冒险家金杯", "冒险家头带"},
	"角斗士的终幕礼": []string{"角斗士的留恋", "角斗士的归宿", "角斗士的希冀", "角斗士的酣醉", "角斗士的凯旋"},
	"祭冰之人":    []string{"", "", "", "", "祭冰礼冠"},
	"千岩牢固":    []string{"勋绩之花", "昭武翎羽", "金铜时晷", "盟誓金爵", "将帅兜鍪"},
	"深林的记忆":   []string{"迷宫的游人", "翠蔓的智者", "贤智的定期", "迷误者之灯", "月桂的宝冠"},
	"勇士之心":    []string{"勇士的勋章", "勇士的期许", "勇士的坚毅", "勇士的壮行", "勇士的冠冕"},
	"赌徒":      []string{"赌徒的胸花", "赌徒的羽饰", "赌徒的怀表", "赌徒的骰盅", "赌徒的耳环"},
	"流放者":     []string{"流放者之花", "流放者之羽", "流放者怀表", "流放者之杯", "流放者头冠"},
	"如雷的盛怒":   []string{"雷鸟的怜悯", "雷灾的孑遗", "雷霆的时计", "降雷的凶兆", "唤雷的头冠"},
	"染血的骑士道":  []string{"染血的铁之心", "染血的黑之羽", "骑士染血之时", "染血骑士之杯", "染血的铁假面"},
	"祭水之人":    []string{"", "", "", "", "祭水礼冠"},
	"追忆之注连":   []string{"羁缠之花", "思忆之矢", "朝露之时", "祈望之心", "无常之面"},
	"冰风迷途的勇士": []string{"历经风雪的思念", "摧冰而行的执望", "冰雪故园的终期", "遍结寒霜的傲骨", "破冰踏雪的回音"},
	"被怜爱的少女":  []string{"远方的少女之心", "少女飘摇的思念", "少女苦短的良辰", "少女片刻的闲暇", "少女易逝的芳颜"},
	"祭火之人":    []string{"", "", "", "", "祭火礼冠"},
	"绝缘之旗印":   []string{"明威之镡", "切落之羽", "雷云之笼", "绯花之壶", "华饰之兜"},
	"饰金之梦":    []string{"梦中的铁花", "裁断的翎羽", "沉金的岁月", "如蜜的终宴", "沙王的投影"},
	"乐园遗落之花":  []string{"月女的华彩", "谢落的筵席", "凝结的时刻", "守秘的魔瓶", "紫晶的花冠"},
}
