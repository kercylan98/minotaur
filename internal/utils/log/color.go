package log

const (
	ColorDefault          = "\033[0m"   // 默认
	ColorDefaultBold      = "\033[1m"   // 默认加粗
	ColorDefaultUnderline = "\033[4m"   // 默认下划线
	ColorDefaultReverse   = "\033[7m"   // 默认反显
	ColorDefaultDelete    = "\033[9m"   // 默认删除线
	ColorBlack            = "\033[30m"  // 黑色
	ColorRed              = "\033[31m"  // 红色
	ColorGreen            = "\033[32m"  // 绿色
	ColorYellow           = "\033[33m"  // 黄色
	ColorBlue             = "\033[34m"  // 蓝色
	ColorPurple           = "\033[35m"  // 紫色
	ColorCyan             = "\033[36m"  // 青色
	ColorWhite            = "\033[37m"  // 白色
	ColorBgBlack          = "\033[40m"  // 背景黑色
	ColorBgRed            = "\033[41m"  // 背景红色
	ColorBgGreen          = "\033[42m"  // 背景绿色
	ColorBgYellow         = "\033[43m"  // 背景黄色
	ColorBgBlue           = "\033[44m"  // 背景蓝色
	ColorBgPurple         = "\033[45m"  // 背景紫色
	ColorBgCyan           = "\033[46m"  // 背景青色
	ColorBgWhite          = "\033[47m"  // 背景白色
	ColorBrightBlack      = "\033[90m"  // 亮黑色
	ColorBrightRed        = "\033[91m"  // 亮红色
	ColorBrightGreen      = "\033[92m"  // 亮绿色
	ColorBrightYellow     = "\033[93m"  // 亮黄色
	ColorBrightBlue       = "\033[94m"  // 亮蓝色
	ColorBrightPurple     = "\033[95m"  // 亮紫色
	ColorBrightCyan       = "\033[96m"  // 亮青色
	ColorBrightWhite      = "\033[97m"  // 亮白色
	ColorBgBrightBlack    = "\033[100m" // 背景亮黑色
	ColorBgBrightRed      = "\033[101m" // 背景亮红色
	ColorBgBrightGreen    = "\033[102m" // 背景亮绿色
	ColorBgBrightYellow   = "\033[103m" // 背景亮黄色
	ColorBgBrightBlue     = "\033[104m" // 背景亮蓝色
	ColorBgBrightPurple   = "\033[105m" // 背景亮紫色
	ColorBgBrightCyan     = "\033[106m" // 背景亮青色
	ColorBgBrightWhite    = "\033[107m" // 背景亮白色

	ColorBlackBold               = "\033[1;30m"  // 黑色加粗
	ColorRedBold                 = "\033[1;31m"  // 红色加粗
	ColorGreenBold               = "\033[1;32m"  // 绿色加粗
	ColorYellowBold              = "\033[1;33m"  // 黄色加粗
	ColorBlueBold                = "\033[1;34m"  // 蓝色加粗
	ColorPurpleBold              = "\033[1;35m"  // 紫色加粗
	ColorCyanBold                = "\033[1;36m"  // 青色加粗
	ColorWhiteBold               = "\033[1;37m"  // 白色加粗
	ColorBgBlackBold             = "\033[1;40m"  // 背景黑色加粗
	ColorBgRedBold               = "\033[1;41m"  // 背景红色加粗
	ColorBgGreenBold             = "\033[1;42m"  // 背景绿色加粗
	ColorBgYellowBold            = "\033[1;43m"  // 背景黄色加粗
	ColorBgBlueBold              = "\033[1;44m"  // 背景蓝色加粗
	ColorBgPurpleBold            = "\033[1;45m"  // 背景紫色加粗
	ColorBgCyanBold              = "\033[1;46m"  // 背景青色加粗
	ColorBgWhiteBold             = "\033[1;47m"  // 背景白色加粗
	ColorBrightBlackBold         = "\033[1;90m"  // 亮黑色加粗
	ColorBrightRedBold           = "\033[1;91m"  // 亮红色加粗
	ColorBrightGreenBold         = "\033[1;92m"  // 亮绿色加粗
	ColorBrightYellowBold        = "\033[1;93m"  // 亮黄色加粗
	ColorBrightBlueBold          = "\033[1;94m"  // 亮蓝色加粗
	ColorBrightPurpleBold        = "\033[1;95m"  // 亮紫色加粗
	ColorBrightCyanBold          = "\033[1;96m"  // 亮青色加粗
	ColorBrightWhiteBold         = "\033[1;97m"  // 亮白色加粗
	ColorBgBrightBlackBold       = "\033[1;100m" // 背景亮黑色加粗
	ColorBgBrightRedBold         = "\033[1;101m" // 背景亮红色加粗
	ColorBgBrightGreenBold       = "\033[1;102m" // 背景亮绿色加粗
	ColorBgBrightYellowBold      = "\033[1;103m" // 背景亮黄色加粗
	ColorBgBrightBlueBold        = "\033[1;104m" // 背景亮蓝色加粗
	ColorBgBrightPurpleBold      = "\033[1;105m" // 背景亮紫色加粗
	ColorBgBrightCyanBold        = "\033[1;106m" // 背景亮青色加粗
	ColorBgBrightWhiteBold       = "\033[1;107m" // 背景亮白色加粗
	ColorBlackUnderline          = "\033[4;30m"  // 黑色下划线
	ColorRedUnderline            = "\033[4;31m"  // 红色下划线
	ColorGreenUnderline          = "\033[4;32m"  // 绿色下划线
	ColorYellowUnderline         = "\033[4;33m"  // 黄色下划线
	ColorBlueUnderline           = "\033[4;34m"  // 蓝色下划线
	ColorPurpleUnderline         = "\033[4;35m"  // 紫色下划线
	ColorCyanUnderline           = "\033[4;36m"  // 青色下划线
	ColorWhiteUnderline          = "\033[4;37m"  // 白色下划线
	ColorBgBlackUnderline        = "\033[4;40m"  // 背景黑色下划线
	ColorBgRedUnderline          = "\033[4;41m"  // 背景红色下划线
	ColorBgGreenUnderline        = "\033[4;42m"  // 背景绿色下划线
	ColorBgYellowUnderline       = "\033[4;43m"  // 背景黄色下划线
	ColorBgBlueUnderline         = "\033[4;44m"  // 背景蓝色下划线
	ColorBgPurpleUnderline       = "\033[4;45m"  // 背景紫色下划线
	ColorBgCyanUnderline         = "\033[4;46m"  // 背景青色下划线
	ColorBgWhiteUnderline        = "\033[4;47m"  // 背景白色下划线
	ColorBrightBlackUnderline    = "\033[4;90m"  // 亮黑色下划线
	ColorBrightRedUnderline      = "\033[4;91m"  // 亮红色下划线
	ColorBrightGreenUnderline    = "\033[4;92m"  // 亮绿色下划线
	ColorBrightYellowUnderline   = "\033[4;93m"  // 亮黄色下划线
	ColorBrightBlueUnderline     = "\033[4;94m"  // 亮蓝色下划线
	ColorBrightPurpleUnderline   = "\033[4;95m"  // 亮紫色下划线
	ColorBrightCyanUnderline     = "\033[4;96m"  // 亮青色下划线
	ColorBrightWhiteUnderline    = "\033[4;97m"  // 亮白色下划线
	ColorBgBrightBlackUnderline  = "\033[4;100m" // 背景亮黑色下划线
	ColorBgBrightRedUnderline    = "\033[4;101m" // 背景亮红色下划线
	ColorBgBrightGreenUnderline  = "\033[4;102m" // 背景亮绿色下划线
	ColorBgBrightYellowUnderline = "\033[4;103m" // 背景亮黄色下划线
	ColorBgBrightBlueUnderline   = "\033[4;104m" // 背景亮蓝色下划线
	ColorBgBrightPurpleUnderline = "\033[4;105m" // 背景亮紫色下划线
	ColorBgBrightCyanUnderline   = "\033[4;106m" // 背景亮青色下划线
	ColorBgBrightWhiteUnderline  = "\033[4;107m" // 背景亮白色下划线
	ColorBlackReverse            = "\033[7;30m"  // 黑色反显
	ColorRedReverse              = "\033[7;31m"  // 红色反显
	ColorGreenReverse            = "\033[7;32m"  // 绿色反显
	ColorYellowReverse           = "\033[7;33m"  // 黄色反显
	ColorBlueReverse             = "\033[7;34m"  // 蓝色反显
	ColorPurpleReverse           = "\033[7;35m"  // 紫色反显
	ColorCyanReverse             = "\033[7;36m"  // 青色反显
	ColorWhiteReverse            = "\033[7;37m"  // 白色反显
	ColorBgBlackReverse          = "\033[7;40m"  // 背景黑色反显
	ColorBgRedReverse            = "\033[7;41m"  // 背景红色反显
	ColorBgGreenReverse          = "\033[7;42m"  // 背景绿色反显
	ColorBgYellowReverse         = "\033[7;43m"  // 背景黄色反显
	ColorBgBlueReverse           = "\033[7;44m"  // 背景蓝色反显
	ColorBgPurpleReverse         = "\033[7;45m"  // 背景紫色反显
	ColorBgCyanReverse           = "\033[7;46m"  // 背景青色反显
	ColorBgWhiteReverse          = "\033[7;47m"  // 背景白色反显
	ColorBrightBlackReverse      = "\033[7;90m"  // 亮黑色反显
	ColorBrightRedReverse        = "\033[7;91m"  // 亮红色反显
	ColorBrightGreenReverse      = "\033[7;92m"  // 亮绿色反显
	ColorBrightYellowReverse     = "\033[7;93m"  // 亮黄色反显
	ColorBrightBlueReverse       = "\033[7;94m"  // 亮蓝色反显
	ColorBrightPurpleReverse     = "\033[7;95m"  // 亮紫色反显
	ColorBrightCyanReverse       = "\033[7;96m"  // 亮青色反显
	ColorBrightWhiteReverse      = "\033[7;97m"  // 亮白色反显
	ColorBgBrightBlackReverse    = "\033[7;100m" // 背景亮黑色反显
	ColorBgBrightRedReverse      = "\033[7;101m" // 背景亮红色反显
	ColorBgBrightGreenReverse    = "\033[7;102m" // 背景亮绿色反显
	ColorBgBrightYellowReverse   = "\033[7;103m" // 背景亮黄色反显
	ColorBgBrightBlueReverse     = "\033[7;104m" // 背景亮蓝色反显
	ColorBgBrightPurpleReverse   = "\033[7;105m" // 背景亮紫色反显
	ColorBgBrightCyanReverse     = "\033[7;106m" // 背景亮青色反显
	ColorBgBrightWhiteReverse    = "\033[7;107m" // 背景亮白色反显
	StrikeThroughBlack           = "\033[9;30m"  // 黑色删除线
	StrikeThroughRed             = "\033[9;31m"  // 红色删除线
	StrikeThroughGreen           = "\033[9;32m"  // 绿色删除线
	StrikeThroughYellow          = "\033[9;33m"  // 黄色删除线
	StrikeThroughBlue            = "\033[9;34m"  // 蓝色删除线
	StrikeThroughPurple          = "\033[9;35m"  // 紫色删除线
	StrikeThroughCyan            = "\033[9;36m"  // 青色删除线
	StrikeThroughWhite           = "\033[9;37m"  // 白色删除线
	BgStrikeThroughBlack         = "\033[9;40m"  // 背景黑色删除线
	BgStrikeThroughRed           = "\033[9;41m"  // 背景红色删除线
	BgStrikeThroughGreen         = "\033[9;42m"  // 背景绿色删除线
	BgStrikeThroughYellow        = "\033[9;43m"  // 背景黄色删除线
	BgStrikeThroughBlue          = "\033[9;44m"  // 背景蓝色删除线
	BgStrikeThroughPurple        = "\033[9;45m"  // 背景紫色删除线
	BgStrikeThroughCyan          = "\033[9;46m"  // 背景青色删除线
	BgStrikeThroughWhite         = "\033[9;47m"  // 背景白色删除线
	BrightStrikeThroughBlack     = "\033[9;90m"  // 亮黑色删除线
	BrightStrikeThroughRed       = "\033[9;91m"  // 亮红色删除线
	BrightStrikeThroughGreen     = "\033[9;92m"  // 亮绿色删除线
	BrightStrikeThroughYellow    = "\033[9;93m"  // 亮黄色删除线
	BrightStrikeThroughBlue      = "\033[9;94m"  // 亮蓝色删除线
	BrightStrikeThroughPurple    = "\033[9;95m"  // 亮紫色删除线
	BrightStrikeThroughCyan      = "\033[9;96m"  // 亮青色删除线
	BrightStrikeThroughWhite     = "\033[9;97m"  // 亮白色删除线
	BgBrightStrikeThroughBlack   = "\033[9;100m" // 背景亮黑色删除线
	BgBrightStrikeThroughRed     = "\033[9;101m" // 背景亮红色删除线
	BgBrightStrikeThroughGreen   = "\033[9;102m" // 背景亮绿色删除线
	BgBrightStrikeThroughYellow  = "\033[9;103m" // 背景亮黄色删除线
	BgBrightStrikeThroughBlue    = "\033[9;104m" // 背景亮蓝色删除线
	BgBrightStrikeThroughPurple  = "\033[9;105m" // 背景亮紫色删除线
	BgBrightStrikeThroughCyan    = "\033[9;106m" // 背景亮青色删除线
	BgBrightStrikeThroughWhite   = "\033[9;107m" // 背景亮白色删除线
)

var colors = map[string]struct{}{
	ColorDefault:                 {},
	ColorDefaultBold:             {},
	ColorDefaultUnderline:        {},
	ColorDefaultReverse:          {},
	ColorDefaultDelete:           {},
	ColorBlack:                   {},
	ColorRed:                     {},
	ColorGreen:                   {},
	ColorYellow:                  {},
	ColorBlue:                    {},
	ColorPurple:                  {},
	ColorCyan:                    {},
	ColorWhite:                   {},
	ColorBgBlack:                 {},
	ColorBgRed:                   {},
	ColorBgGreen:                 {},
	ColorBgYellow:                {},
	ColorBgBlue:                  {},
	ColorBgPurple:                {},
	ColorBgCyan:                  {},
	ColorBgWhite:                 {},
	ColorBrightBlack:             {},
	ColorBrightRed:               {},
	ColorBrightGreen:             {},
	ColorBrightYellow:            {},
	ColorBrightBlue:              {},
	ColorBrightPurple:            {},
	ColorBrightCyan:              {},
	ColorBrightWhite:             {},
	ColorBgBrightBlack:           {},
	ColorBgBrightRed:             {},
	ColorBgBrightGreen:           {},
	ColorBgBrightYellow:          {},
	ColorBgBrightBlue:            {},
	ColorBgBrightPurple:          {},
	ColorBgBrightCyan:            {},
	ColorBgBrightWhite:           {},
	ColorBlackBold:               {},
	ColorRedBold:                 {},
	ColorGreenBold:               {},
	ColorYellowBold:              {},
	ColorBlueBold:                {},
	ColorPurpleBold:              {},
	ColorCyanBold:                {},
	ColorWhiteBold:               {},
	ColorBgBlackBold:             {},
	ColorBgRedBold:               {},
	ColorBgGreenBold:             {},
	ColorBgYellowBold:            {},
	ColorBgBlueBold:              {},
	ColorBgPurpleBold:            {},
	ColorBgCyanBold:              {},
	ColorBgWhiteBold:             {},
	ColorBrightBlackBold:         {},
	ColorBrightRedBold:           {},
	ColorBrightGreenBold:         {},
	ColorBrightYellowBold:        {},
	ColorBrightBlueBold:          {},
	ColorBrightPurpleBold:        {},
	ColorBrightCyanBold:          {},
	ColorBrightWhiteBold:         {},
	ColorBgBrightBlackBold:       {},
	ColorBgBrightRedBold:         {},
	ColorBgBrightGreenBold:       {},
	ColorBgBrightYellowBold:      {},
	ColorBgBrightBlueBold:        {},
	ColorBgBrightPurpleBold:      {},
	ColorBgBrightCyanBold:        {},
	ColorBgBrightWhiteBold:       {},
	ColorBlackUnderline:          {},
	ColorRedUnderline:            {},
	ColorGreenUnderline:          {},
	ColorYellowUnderline:         {},
	ColorBlueUnderline:           {},
	ColorPurpleUnderline:         {},
	ColorCyanUnderline:           {},
	ColorWhiteUnderline:          {},
	ColorBgBlackUnderline:        {},
	ColorBgRedUnderline:          {},
	ColorBgGreenUnderline:        {},
	ColorBgYellowUnderline:       {},
	ColorBgBlueUnderline:         {},
	ColorBgPurpleUnderline:       {},
	ColorBgCyanUnderline:         {},
	ColorBgWhiteUnderline:        {},
	ColorBrightBlackUnderline:    {},
	ColorBrightRedUnderline:      {},
	ColorBrightGreenUnderline:    {},
	ColorBrightYellowUnderline:   {},
	ColorBrightBlueUnderline:     {},
	ColorBrightPurpleUnderline:   {},
	ColorBrightCyanUnderline:     {},
	ColorBrightWhiteUnderline:    {},
	ColorBgBrightBlackUnderline:  {},
	ColorBgBrightRedUnderline:    {},
	ColorBgBrightGreenUnderline:  {},
	ColorBgBrightYellowUnderline: {},
	ColorBgBrightBlueUnderline:   {},
	ColorBgBrightPurpleUnderline: {},
	ColorBgBrightCyanUnderline:   {},
	ColorBgBrightWhiteUnderline:  {},
	ColorBlackReverse:            {},
	ColorRedReverse:              {},
	ColorGreenReverse:            {},
	ColorYellowReverse:           {},
	ColorBlueReverse:             {},
	ColorPurpleReverse:           {},
	ColorCyanReverse:             {},
	ColorWhiteReverse:            {},
	ColorBgBlackReverse:          {},
	ColorBgRedReverse:            {},
	ColorBgGreenReverse:          {},
	ColorBgYellowReverse:         {},
	ColorBgBlueReverse:           {},
	ColorBgPurpleReverse:         {},
	ColorBgCyanReverse:           {},
	ColorBgWhiteReverse:          {},
	ColorBrightBlackReverse:      {},
	ColorBrightRedReverse:        {},
	ColorBrightGreenReverse:      {},
	ColorBrightYellowReverse:     {},
	ColorBrightBlueReverse:       {},
	ColorBrightPurpleReverse:     {},
	ColorBrightCyanReverse:       {},
	ColorBrightWhiteReverse:      {},
	ColorBgBrightBlackReverse:    {},
	ColorBgBrightRedReverse:      {},
	ColorBgBrightGreenReverse:    {},
	ColorBgBrightYellowReverse:   {},
	ColorBgBrightBlueReverse:     {},
	ColorBgBrightPurpleReverse:   {},
	ColorBgBrightCyanReverse:     {},
	ColorBgBrightWhiteReverse:    {},
	StrikeThroughBlack:           {},
	StrikeThroughRed:             {},
	StrikeThroughGreen:           {},
	StrikeThroughYellow:          {},
	StrikeThroughBlue:            {},
	StrikeThroughPurple:          {},
	StrikeThroughCyan:            {},
	StrikeThroughWhite:           {},
	BgStrikeThroughBlack:         {},
	BgStrikeThroughRed:           {},
	BgStrikeThroughGreen:         {},
	BgStrikeThroughYellow:        {},
	BgStrikeThroughBlue:          {},
	BgStrikeThroughPurple:        {},
	BgStrikeThroughCyan:          {},
	BgStrikeThroughWhite:         {},
	BrightStrikeThroughBlack:     {},
	BrightStrikeThroughRed:       {},
	BrightStrikeThroughGreen:     {},
	BrightStrikeThroughYellow:    {},
	BrightStrikeThroughBlue:      {},
	BrightStrikeThroughPurple:    {},
	BrightStrikeThroughCyan:      {},
	BrightStrikeThroughWhite:     {},
	BgBrightStrikeThroughBlack:   {},
	BgBrightStrikeThroughRed:     {},
	BgBrightStrikeThroughGreen:   {},
	BgBrightStrikeThroughYellow:  {},
	BgBrightStrikeThroughBlue:    {},
	BgBrightStrikeThroughPurple:  {},
	BgBrightStrikeThroughCyan:    {},
	BgBrightStrikeThroughWhite:   {},
}
