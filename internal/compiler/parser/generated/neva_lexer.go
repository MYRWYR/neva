// Code generated from ./neva.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type nevaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NevaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nevalexerLexerInit() {
	staticData := &NevaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'import'", "'{'", "'}'", "'@/'", "'.'", "'/'", "'types'", "'<'",
		"'>'", "','", "'enum'", "'['", "']'", "'struct'", "'|'", "'interfaces'",
		"'('", "')'", "'const'", "'true'", "'false'", "'nil'", "':'", "'components'",
		"'#'", "'nodes'", "'net'", "'->'", "'in'", "'out'", "", "'pub'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT", "PUB_KW",
		"IDENTIFIER", "INT", "FLOAT", "STRING", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "COMMENT", "PUB_KW", "IDENTIFIER",
		"LETTER", "INT", "FLOAT", "STRING", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 259, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 203,
		8, 30, 10, 30, 12, 30, 206, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 5, 32, 215, 8, 32, 10, 32, 12, 32, 218, 9, 32, 1, 33, 1, 33,
		1, 34, 4, 34, 223, 8, 34, 11, 34, 12, 34, 224, 1, 35, 5, 35, 228, 8, 35,
		10, 35, 12, 35, 231, 9, 35, 1, 35, 1, 35, 4, 35, 235, 8, 35, 11, 35, 12,
		35, 236, 1, 36, 1, 36, 5, 36, 241, 8, 36, 10, 36, 12, 36, 244, 9, 36, 1,
		36, 1, 36, 1, 37, 3, 37, 249, 8, 37, 1, 37, 1, 37, 1, 38, 4, 38, 254, 8,
		38, 11, 38, 12, 38, 255, 1, 38, 1, 38, 1, 242, 0, 39, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 0, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 1, 0, 4,
		2, 0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0,
		9, 9, 32, 32, 266, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 1, 79, 1, 0, 0, 0, 3, 86, 1, 0, 0, 0, 5, 88, 1, 0, 0, 0, 7, 90,
		1, 0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 95, 1, 0, 0, 0, 13, 97, 1, 0, 0, 0,
		15, 103, 1, 0, 0, 0, 17, 105, 1, 0, 0, 0, 19, 107, 1, 0, 0, 0, 21, 109,
		1, 0, 0, 0, 23, 114, 1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 118, 1, 0, 0,
		0, 29, 125, 1, 0, 0, 0, 31, 127, 1, 0, 0, 0, 33, 138, 1, 0, 0, 0, 35, 140,
		1, 0, 0, 0, 37, 142, 1, 0, 0, 0, 39, 148, 1, 0, 0, 0, 41, 153, 1, 0, 0,
		0, 43, 159, 1, 0, 0, 0, 45, 163, 1, 0, 0, 0, 47, 165, 1, 0, 0, 0, 49, 176,
		1, 0, 0, 0, 51, 178, 1, 0, 0, 0, 53, 184, 1, 0, 0, 0, 55, 188, 1, 0, 0,
		0, 57, 191, 1, 0, 0, 0, 59, 194, 1, 0, 0, 0, 61, 198, 1, 0, 0, 0, 63, 207,
		1, 0, 0, 0, 65, 211, 1, 0, 0, 0, 67, 219, 1, 0, 0, 0, 69, 222, 1, 0, 0,
		0, 71, 229, 1, 0, 0, 0, 73, 238, 1, 0, 0, 0, 75, 248, 1, 0, 0, 0, 77, 253,
		1, 0, 0, 0, 79, 80, 5, 105, 0, 0, 80, 81, 5, 109, 0, 0, 81, 82, 5, 112,
		0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 114, 0, 0, 84, 85, 5, 116, 0, 0,
		85, 2, 1, 0, 0, 0, 86, 87, 5, 123, 0, 0, 87, 4, 1, 0, 0, 0, 88, 89, 5,
		125, 0, 0, 89, 6, 1, 0, 0, 0, 90, 91, 5, 64, 0, 0, 91, 92, 5, 47, 0, 0,
		92, 8, 1, 0, 0, 0, 93, 94, 5, 46, 0, 0, 94, 10, 1, 0, 0, 0, 95, 96, 5,
		47, 0, 0, 96, 12, 1, 0, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5, 121, 0,
		0, 99, 100, 5, 112, 0, 0, 100, 101, 5, 101, 0, 0, 101, 102, 5, 115, 0,
		0, 102, 14, 1, 0, 0, 0, 103, 104, 5, 60, 0, 0, 104, 16, 1, 0, 0, 0, 105,
		106, 5, 62, 0, 0, 106, 18, 1, 0, 0, 0, 107, 108, 5, 44, 0, 0, 108, 20,
		1, 0, 0, 0, 109, 110, 5, 101, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5,
		117, 0, 0, 112, 113, 5, 109, 0, 0, 113, 22, 1, 0, 0, 0, 114, 115, 5, 91,
		0, 0, 115, 24, 1, 0, 0, 0, 116, 117, 5, 93, 0, 0, 117, 26, 1, 0, 0, 0,
		118, 119, 5, 115, 0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 114, 0, 0,
		121, 122, 5, 117, 0, 0, 122, 123, 5, 99, 0, 0, 123, 124, 5, 116, 0, 0,
		124, 28, 1, 0, 0, 0, 125, 126, 5, 124, 0, 0, 126, 30, 1, 0, 0, 0, 127,
		128, 5, 105, 0, 0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 116, 0, 0, 130,
		131, 5, 101, 0, 0, 131, 132, 5, 114, 0, 0, 132, 133, 5, 102, 0, 0, 133,
		134, 5, 97, 0, 0, 134, 135, 5, 99, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137,
		5, 115, 0, 0, 137, 32, 1, 0, 0, 0, 138, 139, 5, 40, 0, 0, 139, 34, 1, 0,
		0, 0, 140, 141, 5, 41, 0, 0, 141, 36, 1, 0, 0, 0, 142, 143, 5, 99, 0, 0,
		143, 144, 5, 111, 0, 0, 144, 145, 5, 110, 0, 0, 145, 146, 5, 115, 0, 0,
		146, 147, 5, 116, 0, 0, 147, 38, 1, 0, 0, 0, 148, 149, 5, 116, 0, 0, 149,
		150, 5, 114, 0, 0, 150, 151, 5, 117, 0, 0, 151, 152, 5, 101, 0, 0, 152,
		40, 1, 0, 0, 0, 153, 154, 5, 102, 0, 0, 154, 155, 5, 97, 0, 0, 155, 156,
		5, 108, 0, 0, 156, 157, 5, 115, 0, 0, 157, 158, 5, 101, 0, 0, 158, 42,
		1, 0, 0, 0, 159, 160, 5, 110, 0, 0, 160, 161, 5, 105, 0, 0, 161, 162, 5,
		108, 0, 0, 162, 44, 1, 0, 0, 0, 163, 164, 5, 58, 0, 0, 164, 46, 1, 0, 0,
		0, 165, 166, 5, 99, 0, 0, 166, 167, 5, 111, 0, 0, 167, 168, 5, 109, 0,
		0, 168, 169, 5, 112, 0, 0, 169, 170, 5, 111, 0, 0, 170, 171, 5, 110, 0,
		0, 171, 172, 5, 101, 0, 0, 172, 173, 5, 110, 0, 0, 173, 174, 5, 116, 0,
		0, 174, 175, 5, 115, 0, 0, 175, 48, 1, 0, 0, 0, 176, 177, 5, 35, 0, 0,
		177, 50, 1, 0, 0, 0, 178, 179, 5, 110, 0, 0, 179, 180, 5, 111, 0, 0, 180,
		181, 5, 100, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 115, 0, 0, 183,
		52, 1, 0, 0, 0, 184, 185, 5, 110, 0, 0, 185, 186, 5, 101, 0, 0, 186, 187,
		5, 116, 0, 0, 187, 54, 1, 0, 0, 0, 188, 189, 5, 45, 0, 0, 189, 190, 5,
		62, 0, 0, 190, 56, 1, 0, 0, 0, 191, 192, 5, 105, 0, 0, 192, 193, 5, 110,
		0, 0, 193, 58, 1, 0, 0, 0, 194, 195, 5, 111, 0, 0, 195, 196, 5, 117, 0,
		0, 196, 197, 5, 116, 0, 0, 197, 60, 1, 0, 0, 0, 198, 199, 5, 47, 0, 0,
		199, 200, 5, 47, 0, 0, 200, 204, 1, 0, 0, 0, 201, 203, 8, 0, 0, 0, 202,
		201, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 62, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 5, 112,
		0, 0, 208, 209, 5, 117, 0, 0, 209, 210, 5, 98, 0, 0, 210, 64, 1, 0, 0,
		0, 211, 216, 3, 67, 33, 0, 212, 215, 3, 67, 33, 0, 213, 215, 3, 69, 34,
		0, 214, 212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 66, 1, 0, 0, 0, 218, 216, 1,
		0, 0, 0, 219, 220, 7, 1, 0, 0, 220, 68, 1, 0, 0, 0, 221, 223, 7, 2, 0,
		0, 222, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 70, 1, 0, 0, 0, 226, 228, 7, 2, 0, 0, 227, 226, 1,
		0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0,
		0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 234, 5, 46, 0, 0, 233,
		235, 7, 2, 0, 0, 234, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 234,
		1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 72, 1, 0, 0, 0, 238, 242, 5, 39,
		0, 0, 239, 241, 9, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0,
		242, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 245, 246, 5, 39, 0, 0, 246, 74, 1, 0, 0, 0, 247, 249,
		5, 13, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0,
		0, 0, 250, 251, 5, 10, 0, 0, 251, 76, 1, 0, 0, 0, 252, 254, 7, 3, 0, 0,
		253, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255,
		256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 6, 38, 0, 0, 258, 78,
		1, 0, 0, 0, 10, 0, 204, 214, 216, 224, 229, 236, 242, 248, 255, 1, 0, 1,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// nevaLexerInit initializes any static state used to implement nevaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewnevaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NevaLexerInit() {
	staticData := &NevaLexerLexerStaticData
	staticData.once.Do(nevalexerLexerInit)
}

// NewnevaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewnevaLexer(input antlr.CharStream) *nevaLexer {
	NevaLexerInit()
	l := new(nevaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NevaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "neva.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// nevaLexer tokens.
const (
	nevaLexerT__0       = 1
	nevaLexerT__1       = 2
	nevaLexerT__2       = 3
	nevaLexerT__3       = 4
	nevaLexerT__4       = 5
	nevaLexerT__5       = 6
	nevaLexerT__6       = 7
	nevaLexerT__7       = 8
	nevaLexerT__8       = 9
	nevaLexerT__9       = 10
	nevaLexerT__10      = 11
	nevaLexerT__11      = 12
	nevaLexerT__12      = 13
	nevaLexerT__13      = 14
	nevaLexerT__14      = 15
	nevaLexerT__15      = 16
	nevaLexerT__16      = 17
	nevaLexerT__17      = 18
	nevaLexerT__18      = 19
	nevaLexerT__19      = 20
	nevaLexerT__20      = 21
	nevaLexerT__21      = 22
	nevaLexerT__22      = 23
	nevaLexerT__23      = 24
	nevaLexerT__24      = 25
	nevaLexerT__25      = 26
	nevaLexerT__26      = 27
	nevaLexerT__27      = 28
	nevaLexerT__28      = 29
	nevaLexerT__29      = 30
	nevaLexerCOMMENT    = 31
	nevaLexerPUB_KW     = 32
	nevaLexerIDENTIFIER = 33
	nevaLexerINT        = 34
	nevaLexerFLOAT      = 35
	nevaLexerSTRING     = 36
	nevaLexerNEWLINE    = 37
	nevaLexerWS         = 38
)
