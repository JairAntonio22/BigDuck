// Code generated from BigDuck.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // BigDuck

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 386,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	110, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 123, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 129, 10, 6, 3, 7,
	3, 7, 5, 7, 133, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 149, 10, 11, 3, 12,
	3, 12, 3, 12, 5, 12, 154, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 159, 10,
	13, 3, 13, 3, 13, 5, 13, 163, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 178, 10,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	189, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 195, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 5, 20, 206, 10, 20,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5, 22, 214, 10, 22, 3, 23, 3,
	23, 5, 23, 218, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 232, 10, 24, 3, 24, 5, 24, 235,
	10, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 251, 10, 29, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 5, 31, 259, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 268, 10, 32, 3, 32, 3, 32, 3, 32, 5, 32, 273,
	10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 279, 10, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 291, 10, 34,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 298, 10, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 308, 10, 37, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 5, 38, 324, 10, 38, 3, 39, 3, 39, 3, 39, 5, 39, 329, 10, 39,
	3, 39, 3, 39, 3, 39, 5, 39, 334, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 41, 3, 42, 3, 42, 5, 42, 345, 10, 42, 3, 43, 3, 43, 3, 43,
	5, 43, 350, 10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 356, 10, 44, 3,
	44, 3, 44, 3, 45, 3, 45, 5, 45, 362, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3,
	50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 384, 10, 51, 3, 51, 2, 2,
	52, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 2, 7, 3, 2, 34,
	36, 3, 2, 10, 15, 3, 2, 16, 17, 3, 2, 18, 19, 3, 2, 28, 29, 2, 382, 2,
	102, 3, 2, 2, 2, 4, 109, 3, 2, 2, 2, 6, 111, 3, 2, 2, 2, 8, 122, 3, 2,
	2, 2, 10, 128, 3, 2, 2, 2, 12, 132, 3, 2, 2, 2, 14, 134, 3, 2, 2, 2, 16,
	136, 3, 2, 2, 2, 18, 139, 3, 2, 2, 2, 20, 148, 3, 2, 2, 2, 22, 150, 3,
	2, 2, 2, 24, 155, 3, 2, 2, 2, 26, 166, 3, 2, 2, 2, 28, 170, 3, 2, 2, 2,
	30, 188, 3, 2, 2, 2, 32, 194, 3, 2, 2, 2, 34, 196, 3, 2, 2, 2, 36, 199,
	3, 2, 2, 2, 38, 205, 3, 2, 2, 2, 40, 207, 3, 2, 2, 2, 42, 213, 3, 2, 2,
	2, 44, 217, 3, 2, 2, 2, 46, 234, 3, 2, 2, 2, 48, 236, 3, 2, 2, 2, 50, 239,
	3, 2, 2, 2, 52, 242, 3, 2, 2, 2, 54, 244, 3, 2, 2, 2, 56, 250, 3, 2, 2,
	2, 58, 252, 3, 2, 2, 2, 60, 258, 3, 2, 2, 2, 62, 272, 3, 2, 2, 2, 64, 274,
	3, 2, 2, 2, 66, 290, 3, 2, 2, 2, 68, 297, 3, 2, 2, 2, 70, 299, 3, 2, 2,
	2, 72, 307, 3, 2, 2, 2, 74, 323, 3, 2, 2, 2, 76, 325, 3, 2, 2, 2, 78, 335,
	3, 2, 2, 2, 80, 339, 3, 2, 2, 2, 82, 344, 3, 2, 2, 2, 84, 346, 3, 2, 2,
	2, 86, 351, 3, 2, 2, 2, 88, 361, 3, 2, 2, 2, 90, 368, 3, 2, 2, 2, 92, 370,
	3, 2, 2, 2, 94, 372, 3, 2, 2, 2, 96, 374, 3, 2, 2, 2, 98, 376, 3, 2, 2,
	2, 100, 378, 3, 2, 2, 2, 102, 103, 5, 4, 3, 2, 103, 104, 5, 22, 12, 2,
	104, 3, 3, 2, 2, 2, 105, 106, 5, 6, 4, 2, 106, 107, 5, 4, 3, 2, 107, 110,
	3, 2, 2, 2, 108, 110, 3, 2, 2, 2, 109, 105, 3, 2, 2, 2, 109, 108, 3, 2,
	2, 2, 110, 5, 3, 2, 2, 2, 111, 112, 7, 33, 2, 2, 112, 113, 7, 42, 2, 2,
	113, 114, 5, 8, 5, 2, 114, 115, 5, 12, 7, 2, 115, 116, 7, 3, 2, 2, 116,
	117, 5, 10, 6, 2, 117, 7, 3, 2, 2, 2, 118, 119, 7, 4, 2, 2, 119, 120, 7,
	42, 2, 2, 120, 123, 5, 8, 5, 2, 121, 123, 3, 2, 2, 2, 122, 118, 3, 2, 2,
	2, 122, 121, 3, 2, 2, 2, 123, 9, 3, 2, 2, 2, 124, 125, 5, 6, 4, 2, 125,
	126, 5, 10, 6, 2, 126, 129, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 124,
	3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 11, 3, 2, 2, 2, 130, 133, 5, 14,
	8, 2, 131, 133, 5, 16, 9, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2,
	133, 13, 3, 2, 2, 2, 134, 135, 9, 2, 2, 2, 135, 15, 3, 2, 2, 2, 136, 137,
	5, 18, 10, 2, 137, 138, 5, 14, 8, 2, 138, 17, 3, 2, 2, 2, 139, 140, 7,
	5, 2, 2, 140, 141, 5, 54, 28, 2, 141, 142, 7, 6, 2, 2, 142, 143, 5, 20,
	11, 2, 143, 19, 3, 2, 2, 2, 144, 145, 5, 18, 10, 2, 145, 146, 5, 20, 11,
	2, 146, 149, 3, 2, 2, 2, 147, 149, 3, 2, 2, 2, 148, 144, 3, 2, 2, 2, 148,
	147, 3, 2, 2, 2, 149, 21, 3, 2, 2, 2, 150, 153, 5, 24, 13, 2, 151, 154,
	5, 22, 12, 2, 152, 154, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 152, 3,
	2, 2, 2, 154, 23, 3, 2, 2, 2, 155, 158, 5, 26, 14, 2, 156, 159, 5, 34,
	18, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2, 2, 2,
	159, 162, 3, 2, 2, 2, 160, 163, 5, 6, 4, 2, 161, 163, 3, 2, 2, 2, 162,
	160, 3, 2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165,
	5, 70, 36, 2, 165, 25, 3, 2, 2, 2, 166, 167, 7, 23, 2, 2, 167, 168, 7,
	42, 2, 2, 168, 169, 5, 28, 15, 2, 169, 27, 3, 2, 2, 2, 170, 177, 7, 7,
	2, 2, 171, 172, 7, 42, 2, 2, 172, 173, 5, 32, 17, 2, 173, 174, 5, 14, 8,
	2, 174, 175, 5, 30, 16, 2, 175, 178, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2,
	177, 171, 3, 2, 2, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	180, 7, 8, 2, 2, 180, 29, 3, 2, 2, 2, 181, 182, 7, 3, 2, 2, 182, 183, 7,
	42, 2, 2, 183, 184, 5, 32, 17, 2, 184, 185, 5, 12, 7, 2, 185, 186, 5, 30,
	16, 2, 186, 189, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 181, 3, 2, 2, 2,
	188, 187, 3, 2, 2, 2, 189, 31, 3, 2, 2, 2, 190, 191, 7, 4, 2, 2, 191, 192,
	7, 42, 2, 2, 192, 195, 5, 32, 17, 2, 193, 195, 3, 2, 2, 2, 194, 190, 3,
	2, 2, 2, 194, 193, 3, 2, 2, 2, 195, 33, 3, 2, 2, 2, 196, 197, 7, 9, 2,
	2, 197, 198, 5, 14, 8, 2, 198, 35, 3, 2, 2, 2, 199, 200, 5, 40, 21, 2,
	200, 201, 5, 38, 20, 2, 201, 37, 3, 2, 2, 2, 202, 203, 7, 31, 2, 2, 203,
	206, 5, 36, 19, 2, 204, 206, 3, 2, 2, 2, 205, 202, 3, 2, 2, 2, 205, 204,
	3, 2, 2, 2, 206, 39, 3, 2, 2, 2, 207, 208, 5, 44, 23, 2, 208, 209, 5, 42,
	22, 2, 209, 41, 3, 2, 2, 2, 210, 211, 7, 30, 2, 2, 211, 214, 5, 40, 21,
	2, 212, 214, 3, 2, 2, 2, 213, 210, 3, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214,
	43, 3, 2, 2, 2, 215, 218, 7, 32, 2, 2, 216, 218, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 5, 46,
	24, 2, 220, 45, 3, 2, 2, 2, 221, 222, 7, 7, 2, 2, 222, 223, 5, 36, 19,
	2, 223, 224, 7, 8, 2, 2, 224, 235, 3, 2, 2, 2, 225, 235, 5, 48, 25, 2,
	226, 235, 7, 37, 2, 2, 227, 235, 7, 38, 2, 2, 228, 231, 7, 42, 2, 2, 229,
	232, 5, 18, 10, 2, 230, 232, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 230,
	3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 235, 5, 64, 33, 2, 234, 221, 3,
	2, 2, 2, 234, 225, 3, 2, 2, 2, 234, 226, 3, 2, 2, 2, 234, 227, 3, 2, 2,
	2, 234, 228, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 47, 3, 2, 2, 2, 236,
	237, 5, 54, 28, 2, 237, 238, 5, 50, 26, 2, 238, 49, 3, 2, 2, 2, 239, 240,
	5, 52, 27, 2, 240, 241, 5, 54, 28, 2, 241, 51, 3, 2, 2, 2, 242, 243, 9,
	3, 2, 2, 243, 53, 3, 2, 2, 2, 244, 245, 5, 58, 30, 2, 245, 246, 5, 56,
	29, 2, 246, 55, 3, 2, 2, 2, 247, 248, 9, 4, 2, 2, 248, 251, 5, 54, 28,
	2, 249, 251, 3, 2, 2, 2, 250, 247, 3, 2, 2, 2, 250, 249, 3, 2, 2, 2, 251,
	57, 3, 2, 2, 2, 252, 253, 5, 62, 32, 2, 253, 254, 5, 60, 31, 2, 254, 59,
	3, 2, 2, 2, 255, 256, 9, 5, 2, 2, 256, 259, 5, 58, 30, 2, 257, 259, 3,
	2, 2, 2, 258, 255, 3, 2, 2, 2, 258, 257, 3, 2, 2, 2, 259, 61, 3, 2, 2,
	2, 260, 261, 7, 7, 2, 2, 261, 262, 5, 54, 28, 2, 262, 263, 7, 8, 2, 2,
	263, 273, 3, 2, 2, 2, 264, 267, 7, 42, 2, 2, 265, 268, 5, 18, 10, 2, 266,
	268, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 273,
	3, 2, 2, 2, 269, 273, 7, 39, 2, 2, 270, 273, 7, 40, 2, 2, 271, 273, 5,
	64, 33, 2, 272, 260, 3, 2, 2, 2, 272, 264, 3, 2, 2, 2, 272, 269, 3, 2,
	2, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 2, 2, 2, 273, 63, 3, 2, 2, 2,
	274, 275, 7, 42, 2, 2, 275, 278, 7, 7, 2, 2, 276, 279, 5, 66, 34, 2, 277,
	279, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 277, 3, 2, 2, 2, 279, 280,
	3, 2, 2, 2, 280, 281, 7, 8, 2, 2, 281, 65, 3, 2, 2, 2, 282, 283, 5, 36,
	19, 2, 283, 284, 5, 68, 35, 2, 284, 291, 3, 2, 2, 2, 285, 286, 5, 54, 28,
	2, 286, 287, 5, 68, 35, 2, 287, 291, 3, 2, 2, 2, 288, 289, 7, 41, 2, 2,
	289, 291, 5, 68, 35, 2, 290, 282, 3, 2, 2, 2, 290, 285, 3, 2, 2, 2, 290,
	288, 3, 2, 2, 2, 291, 67, 3, 2, 2, 2, 292, 293, 7, 4, 2, 2, 293, 294, 5,
	66, 34, 2, 294, 295, 5, 68, 35, 2, 295, 298, 3, 2, 2, 2, 296, 298, 3, 2,
	2, 2, 297, 292, 3, 2, 2, 2, 297, 296, 3, 2, 2, 2, 298, 69, 3, 2, 2, 2,
	299, 300, 7, 20, 2, 2, 300, 301, 5, 72, 37, 2, 301, 302, 7, 21, 2, 2, 302,
	71, 3, 2, 2, 2, 303, 304, 5, 74, 38, 2, 304, 305, 5, 72, 37, 2, 305, 308,
	3, 2, 2, 2, 306, 308, 3, 2, 2, 2, 307, 303, 3, 2, 2, 2, 307, 306, 3, 2,
	2, 2, 308, 73, 3, 2, 2, 2, 309, 310, 5, 76, 39, 2, 310, 311, 7, 3, 2, 2,
	311, 324, 3, 2, 2, 2, 312, 324, 5, 78, 40, 2, 313, 324, 5, 86, 44, 2, 314,
	315, 5, 98, 50, 2, 315, 316, 7, 3, 2, 2, 316, 324, 3, 2, 2, 2, 317, 318,
	5, 100, 51, 2, 318, 319, 7, 3, 2, 2, 319, 324, 3, 2, 2, 2, 320, 321, 5,
	64, 33, 2, 321, 322, 7, 3, 2, 2, 322, 324, 3, 2, 2, 2, 323, 309, 3, 2,
	2, 2, 323, 312, 3, 2, 2, 2, 323, 313, 3, 2, 2, 2, 323, 314, 3, 2, 2, 2,
	323, 317, 3, 2, 2, 2, 323, 320, 3, 2, 2, 2, 324, 75, 3, 2, 2, 2, 325, 328,
	7, 42, 2, 2, 326, 329, 5, 18, 10, 2, 327, 329, 3, 2, 2, 2, 328, 326, 3,
	2, 2, 2, 328, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 333, 7, 22, 2,
	2, 331, 334, 5, 54, 28, 2, 332, 334, 5, 36, 19, 2, 333, 331, 3, 2, 2, 2,
	333, 332, 3, 2, 2, 2, 334, 77, 3, 2, 2, 2, 335, 336, 7, 25, 2, 2, 336,
	337, 5, 36, 19, 2, 337, 338, 5, 80, 41, 2, 338, 79, 3, 2, 2, 2, 339, 340,
	5, 70, 36, 2, 340, 341, 5, 82, 42, 2, 341, 81, 3, 2, 2, 2, 342, 345, 5,
	84, 43, 2, 343, 345, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 344, 343, 3, 2,
	2, 2, 345, 83, 3, 2, 2, 2, 346, 349, 7, 26, 2, 2, 347, 350, 5, 78, 40,
	2, 348, 350, 5, 70, 36, 2, 349, 347, 3, 2, 2, 2, 349, 348, 3, 2, 2, 2,
	350, 85, 3, 2, 2, 2, 351, 355, 7, 27, 2, 2, 352, 356, 5, 88, 45, 2, 353,
	356, 5, 94, 48, 2, 354, 356, 5, 96, 49, 2, 355, 352, 3, 2, 2, 2, 355, 353,
	3, 2, 2, 2, 355, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 5, 70,
	36, 2, 358, 87, 3, 2, 2, 2, 359, 362, 5, 76, 39, 2, 360, 362, 3, 2, 2,
	2, 361, 359, 3, 2, 2, 2, 361, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363,
	364, 7, 3, 2, 2, 364, 365, 5, 90, 46, 2, 365, 366, 7, 3, 2, 2, 366, 367,
	5, 92, 47, 2, 367, 89, 3, 2, 2, 2, 368, 369, 5, 36, 19, 2, 369, 91, 3,
	2, 2, 2, 370, 371, 5, 76, 39, 2, 371, 93, 3, 2, 2, 2, 372, 373, 5, 36,
	19, 2, 373, 95, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 97, 3, 2, 2, 2,
	376, 377, 9, 6, 2, 2, 377, 99, 3, 2, 2, 2, 378, 383, 7, 24, 2, 2, 379,
	384, 5, 54, 28, 2, 380, 384, 5, 36, 19, 2, 381, 384, 5, 64, 33, 2, 382,
	384, 3, 2, 2, 2, 383, 379, 3, 2, 2, 2, 383, 380, 3, 2, 2, 2, 383, 381,
	3, 2, 2, 2, 383, 382, 3, 2, 2, 2, 384, 101, 3, 2, 2, 2, 34, 109, 122, 128,
	132, 148, 153, 158, 162, 177, 188, 194, 205, 213, 217, 231, 234, 250, 258,
	267, 272, 278, 290, 297, 307, 323, 328, 333, 344, 349, 355, 361, 383,
}
var literalNames = []string{
	"", "';'", "','", "'['", "']'", "'('", "')'", "'->'", "'='", "'/='", "'<'",
	"'>'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'{'", "'}'", "'<-'",
	"'proc'", "'return'", "'if'", "'else'", "'loop'", "'break'", "'skip'",
	"'and'", "'or'", "'not'", "'var'", "'int'", "'float'", "'bool'", "'true'",
	"'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "PROC", "RETURN", "IF", "ELSE", "LOOP", "BREAK", "SKIP_W",
	"AND", "OR", "NOT", "VAR", "INT", "FLOAT", "BOOL", "TRUE", "FALSE", "CTE_INT",
	"CTE_FLOAT", "CTE_STRING", "ID", "WS", "COMMENT",
}

var ruleNames = []string{
	"program", "vars_decl", "var_decl", "nextVar", "nextVarDecl", "var_type",
	"scalar", "tensor", "dim", "nextDim", "procs_decl", "proc_decl", "sign",
	"args", "nextTypes", "nextArg", "ret_type", "bool_expr", "nextBool", "and_expr",
	"nextAnd", "not_expr", "bool_term", "rel_expr", "opRel", "relOp", "num_expr",
	"nextSum", "prod_expr", "nextProd", "factor", "proc_call", "param", "nextParam",
	"block", "stmts", "stmt", "assignment", "condition", "bodyCond", "endIfBlock",
	"alter", "loop_stmt", "forStyle", "forCond", "ctrl_var", "whileStyle",
	"infLoop", "ctrl_flow", "ret_stmt",
}

type BigDuckParser struct {
	*antlr.BaseParser
}

// NewBigDuckParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BigDuckParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBigDuckParser(input antlr.TokenStream) *BigDuckParser {
	this := new(BigDuckParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BigDuck.g4"

	return this
}

// BigDuckParser tokens.
const (
	BigDuckParserEOF        = antlr.TokenEOF
	BigDuckParserT__0       = 1
	BigDuckParserT__1       = 2
	BigDuckParserT__2       = 3
	BigDuckParserT__3       = 4
	BigDuckParserT__4       = 5
	BigDuckParserT__5       = 6
	BigDuckParserT__6       = 7
	BigDuckParserT__7       = 8
	BigDuckParserT__8       = 9
	BigDuckParserT__9       = 10
	BigDuckParserT__10      = 11
	BigDuckParserT__11      = 12
	BigDuckParserT__12      = 13
	BigDuckParserT__13      = 14
	BigDuckParserT__14      = 15
	BigDuckParserT__15      = 16
	BigDuckParserT__16      = 17
	BigDuckParserT__17      = 18
	BigDuckParserT__18      = 19
	BigDuckParserT__19      = 20
	BigDuckParserPROC       = 21
	BigDuckParserRETURN     = 22
	BigDuckParserIF         = 23
	BigDuckParserELSE       = 24
	BigDuckParserLOOP       = 25
	BigDuckParserBREAK      = 26
	BigDuckParserSKIP_W     = 27
	BigDuckParserAND        = 28
	BigDuckParserOR         = 29
	BigDuckParserNOT        = 30
	BigDuckParserVAR        = 31
	BigDuckParserINT        = 32
	BigDuckParserFLOAT      = 33
	BigDuckParserBOOL       = 34
	BigDuckParserTRUE       = 35
	BigDuckParserFALSE      = 36
	BigDuckParserCTE_INT    = 37
	BigDuckParserCTE_FLOAT  = 38
	BigDuckParserCTE_STRING = 39
	BigDuckParserID         = 40
	BigDuckParserWS         = 41
	BigDuckParserCOMMENT    = 42
)

// BigDuckParser rules.
const (
	BigDuckParserRULE_program     = 0
	BigDuckParserRULE_vars_decl   = 1
	BigDuckParserRULE_var_decl    = 2
	BigDuckParserRULE_nextVar     = 3
	BigDuckParserRULE_nextVarDecl = 4
	BigDuckParserRULE_var_type    = 5
	BigDuckParserRULE_scalar      = 6
	BigDuckParserRULE_tensor      = 7
	BigDuckParserRULE_dim         = 8
	BigDuckParserRULE_nextDim     = 9
	BigDuckParserRULE_procs_decl  = 10
	BigDuckParserRULE_proc_decl   = 11
	BigDuckParserRULE_sign        = 12
	BigDuckParserRULE_args        = 13
	BigDuckParserRULE_nextTypes   = 14
	BigDuckParserRULE_nextArg     = 15
	BigDuckParserRULE_ret_type    = 16
	BigDuckParserRULE_bool_expr   = 17
	BigDuckParserRULE_nextBool    = 18
	BigDuckParserRULE_and_expr    = 19
	BigDuckParserRULE_nextAnd     = 20
	BigDuckParserRULE_not_expr    = 21
	BigDuckParserRULE_bool_term   = 22
	BigDuckParserRULE_rel_expr    = 23
	BigDuckParserRULE_opRel       = 24
	BigDuckParserRULE_relOp       = 25
	BigDuckParserRULE_num_expr    = 26
	BigDuckParserRULE_nextSum     = 27
	BigDuckParserRULE_prod_expr   = 28
	BigDuckParserRULE_nextProd    = 29
	BigDuckParserRULE_factor      = 30
	BigDuckParserRULE_proc_call   = 31
	BigDuckParserRULE_param       = 32
	BigDuckParserRULE_nextParam   = 33
	BigDuckParserRULE_block       = 34
	BigDuckParserRULE_stmts       = 35
	BigDuckParserRULE_stmt        = 36
	BigDuckParserRULE_assignment  = 37
	BigDuckParserRULE_condition   = 38
	BigDuckParserRULE_bodyCond    = 39
	BigDuckParserRULE_endIfBlock  = 40
	BigDuckParserRULE_alter       = 41
	BigDuckParserRULE_loop_stmt   = 42
	BigDuckParserRULE_forStyle    = 43
	BigDuckParserRULE_forCond     = 44
	BigDuckParserRULE_ctrl_var    = 45
	BigDuckParserRULE_whileStyle  = 46
	BigDuckParserRULE_infLoop     = 47
	BigDuckParserRULE_ctrl_flow   = 48
	BigDuckParserRULE_ret_stmt    = 49
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Vars_decl() IVars_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_declContext)
}

func (s *ProgramContext) Procs_decl() IProcs_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcs_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcs_declContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BigDuckParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BigDuckParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Vars_decl()
	}
	{
		p.SetState(101)
		p.Procs_decl()
	}

	return localctx
}

// IVars_declContext is an interface to support dynamic dispatch.
type IVars_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_declContext differentiates from other interfaces.
	IsVars_declContext()
}

type Vars_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_declContext() *Vars_declContext {
	var p = new(Vars_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_vars_decl
	return p
}

func (*Vars_declContext) IsVars_declContext() {}

func NewVars_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_declContext {
	var p = new(Vars_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_vars_decl

	return p
}

func (s *Vars_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_declContext) Var_decl() IVar_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Vars_declContext) Vars_decl() IVars_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_declContext)
}

func (s *Vars_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterVars_decl(s)
	}
}

func (s *Vars_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitVars_decl(s)
	}
}

func (p *BigDuckParser) Vars_decl() (localctx IVars_declContext) {
	localctx = NewVars_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BigDuckParserRULE_vars_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Var_decl()
		}
		{
			p.SetState(104)
			p.Vars_decl()
		}

	case BigDuckParserPROC:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_var_decl
	return p
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) VAR() antlr.TerminalNode {
	return s.GetToken(BigDuckParserVAR, 0)
}

func (s *Var_declContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *Var_declContext) NextVar() INextVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextVarContext)
}

func (s *Var_declContext) Var_type() IVar_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Var_declContext) NextVarDecl() INextVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextVarDeclContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (p *BigDuckParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BigDuckParserRULE_var_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(BigDuckParserVAR)
	}
	{
		p.SetState(110)
		p.Match(BigDuckParserID)
	}
	{
		p.SetState(111)
		p.NextVar()
	}
	{
		p.SetState(112)
		p.Var_type()
	}
	{
		p.SetState(113)
		p.Match(BigDuckParserT__0)
	}
	{
		p.SetState(114)
		p.NextVarDecl()
	}

	return localctx
}

// INextVarContext is an interface to support dynamic dispatch.
type INextVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextVarContext differentiates from other interfaces.
	IsNextVarContext()
}

type NextVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextVarContext() *NextVarContext {
	var p = new(NextVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextVar
	return p
}

func (*NextVarContext) IsNextVarContext() {}

func NewNextVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextVarContext {
	var p = new(NextVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextVar

	return p
}

func (s *NextVarContext) GetParser() antlr.Parser { return s.parser }

func (s *NextVarContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *NextVarContext) NextVar() INextVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextVarContext)
}

func (s *NextVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextVar(s)
	}
}

func (s *NextVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextVar(s)
	}
}

func (p *BigDuckParser) NextVar() (localctx INextVarContext) {
	localctx = NewNextVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BigDuckParserRULE_nextVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(BigDuckParserT__1)
		}
		{
			p.SetState(117)
			p.Match(BigDuckParserID)
		}
		{
			p.SetState(118)
			p.NextVar()
		}

	case BigDuckParserT__2, BigDuckParserINT, BigDuckParserFLOAT, BigDuckParserBOOL:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INextVarDeclContext is an interface to support dynamic dispatch.
type INextVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextVarDeclContext differentiates from other interfaces.
	IsNextVarDeclContext()
}

type NextVarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextVarDeclContext() *NextVarDeclContext {
	var p = new(NextVarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextVarDecl
	return p
}

func (*NextVarDeclContext) IsNextVarDeclContext() {}

func NewNextVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextVarDeclContext {
	var p = new(NextVarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextVarDecl

	return p
}

func (s *NextVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NextVarDeclContext) Var_decl() IVar_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *NextVarDeclContext) NextVarDecl() INextVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextVarDeclContext)
}

func (s *NextVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextVarDecl(s)
	}
}

func (s *NextVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextVarDecl(s)
	}
}

func (p *BigDuckParser) NextVarDecl() (localctx INextVarDeclContext) {
	localctx = NewNextVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BigDuckParserRULE_nextVarDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Var_decl()
		}
		{
			p.SetState(123)
			p.NextVarDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_var_type
	return p
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *Var_typeContext) Tensor() ITensorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITensorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITensorContext)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *BigDuckParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BigDuckParserRULE_var_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserINT, BigDuckParserFLOAT, BigDuckParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Scalar()
		}

	case BigDuckParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Tensor()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IScalarContext is an interface to support dynamic dispatch.
type IScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarContext differentiates from other interfaces.
	IsScalarContext()
}

type ScalarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarContext() *ScalarContext {
	var p = new(ScalarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_scalar
	return p
}

func (*ScalarContext) IsScalarContext() {}

func NewScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarContext {
	var p = new(ScalarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_scalar

	return p
}

func (s *ScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarContext) INT() antlr.TerminalNode {
	return s.GetToken(BigDuckParserINT, 0)
}

func (s *ScalarContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BigDuckParserFLOAT, 0)
}

func (s *ScalarContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BigDuckParserBOOL, 0)
}

func (s *ScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterScalar(s)
	}
}

func (s *ScalarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitScalar(s)
	}
}

func (p *BigDuckParser) Scalar() (localctx IScalarContext) {
	localctx = NewScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BigDuckParserRULE_scalar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BigDuckParserINT-32))|(1<<(BigDuckParserFLOAT-32))|(1<<(BigDuckParserBOOL-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITensorContext is an interface to support dynamic dispatch.
type ITensorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTensorContext differentiates from other interfaces.
	IsTensorContext()
}

type TensorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTensorContext() *TensorContext {
	var p = new(TensorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_tensor
	return p
}

func (*TensorContext) IsTensorContext() {}

func NewTensorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TensorContext {
	var p = new(TensorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_tensor

	return p
}

func (s *TensorContext) GetParser() antlr.Parser { return s.parser }

func (s *TensorContext) Dim() IDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimContext)
}

func (s *TensorContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *TensorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TensorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TensorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterTensor(s)
	}
}

func (s *TensorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitTensor(s)
	}
}

func (p *BigDuckParser) Tensor() (localctx ITensorContext) {
	localctx = NewTensorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BigDuckParserRULE_tensor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Dim()
	}
	{
		p.SetState(135)
		p.Scalar()
	}

	return localctx
}

// IDimContext is an interface to support dynamic dispatch.
type IDimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDimContext differentiates from other interfaces.
	IsDimContext()
}

type DimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimContext() *DimContext {
	var p = new(DimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_dim
	return p
}

func (*DimContext) IsDimContext() {}

func NewDimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimContext {
	var p = new(DimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_dim

	return p
}

func (s *DimContext) GetParser() antlr.Parser { return s.parser }

func (s *DimContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *DimContext) NextDim() INextDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextDimContext)
}

func (s *DimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterDim(s)
	}
}

func (s *DimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitDim(s)
	}
}

func (p *BigDuckParser) Dim() (localctx IDimContext) {
	localctx = NewDimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BigDuckParserRULE_dim)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(BigDuckParserT__2)
	}
	{
		p.SetState(138)
		p.Num_expr()
	}
	{
		p.SetState(139)
		p.Match(BigDuckParserT__3)
	}
	{
		p.SetState(140)
		p.NextDim()
	}

	return localctx
}

// INextDimContext is an interface to support dynamic dispatch.
type INextDimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextDimContext differentiates from other interfaces.
	IsNextDimContext()
}

type NextDimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextDimContext() *NextDimContext {
	var p = new(NextDimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextDim
	return p
}

func (*NextDimContext) IsNextDimContext() {}

func NewNextDimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextDimContext {
	var p = new(NextDimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextDim

	return p
}

func (s *NextDimContext) GetParser() antlr.Parser { return s.parser }

func (s *NextDimContext) Dim() IDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimContext)
}

func (s *NextDimContext) NextDim() INextDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextDimContext)
}

func (s *NextDimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextDimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextDimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextDim(s)
	}
}

func (s *NextDimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextDim(s)
	}
}

func (p *BigDuckParser) NextDim() (localctx INextDimContext) {
	localctx = NewNextDimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BigDuckParserRULE_nextDim)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Dim()
		}
		{
			p.SetState(143)
			p.NextDim()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// IProcs_declContext is an interface to support dynamic dispatch.
type IProcs_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcs_declContext differentiates from other interfaces.
	IsProcs_declContext()
}

type Procs_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcs_declContext() *Procs_declContext {
	var p = new(Procs_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_procs_decl
	return p
}

func (*Procs_declContext) IsProcs_declContext() {}

func NewProcs_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procs_declContext {
	var p = new(Procs_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_procs_decl

	return p
}

func (s *Procs_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Procs_declContext) Proc_decl() IProc_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProc_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProc_declContext)
}

func (s *Procs_declContext) Procs_decl() IProcs_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcs_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcs_declContext)
}

func (s *Procs_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procs_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Procs_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterProcs_decl(s)
	}
}

func (s *Procs_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitProcs_decl(s)
	}
}

func (p *BigDuckParser) Procs_decl() (localctx IProcs_declContext) {
	localctx = NewProcs_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BigDuckParserRULE_procs_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Proc_decl()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserPROC:
		{
			p.SetState(149)
			p.Procs_decl()
		}

	case BigDuckParserEOF:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProc_declContext is an interface to support dynamic dispatch.
type IProc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProc_declContext differentiates from other interfaces.
	IsProc_declContext()
}

type Proc_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_declContext() *Proc_declContext {
	var p = new(Proc_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_proc_decl
	return p
}

func (*Proc_declContext) IsProc_declContext() {}

func NewProc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_declContext {
	var p = new(Proc_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_proc_decl

	return p
}

func (s *Proc_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_declContext) Sign() ISignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *Proc_declContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Proc_declContext) Ret_type() IRet_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRet_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRet_typeContext)
}

func (s *Proc_declContext) Var_decl() IVar_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Proc_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterProc_decl(s)
	}
}

func (s *Proc_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitProc_decl(s)
	}
}

func (p *BigDuckParser) Proc_decl() (localctx IProc_declContext) {
	localctx = NewProc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BigDuckParserRULE_proc_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Sign()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__6:
		{
			p.SetState(154)
			p.Ret_type()
		}

	case BigDuckParserT__17, BigDuckParserVAR:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserVAR:
		{
			p.SetState(158)
			p.Var_decl()
		}

	case BigDuckParserT__17:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(162)
		p.Block()
	}

	return localctx
}

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_sign
	return p
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) PROC() antlr.TerminalNode {
	return s.GetToken(BigDuckParserPROC, 0)
}

func (s *SignContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *SignContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterSign(s)
	}
}

func (s *SignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitSign(s)
	}
}

func (p *BigDuckParser) Sign() (localctx ISignContext) {
	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BigDuckParserRULE_sign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(BigDuckParserPROC)
	}
	{
		p.SetState(165)
		p.Match(BigDuckParserID)
	}
	{
		p.SetState(166)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *ArgsContext) NextArg() INextArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextArgContext)
}

func (s *ArgsContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *ArgsContext) NextTypes() INextTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextTypesContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *BigDuckParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BigDuckParserRULE_args)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(BigDuckParserT__4)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserID:
		{
			p.SetState(169)
			p.Match(BigDuckParserID)
		}
		{
			p.SetState(170)
			p.NextArg()
		}
		{
			p.SetState(171)
			p.Scalar()
		}
		{
			p.SetState(172)
			p.NextTypes()
		}

	case BigDuckParserT__5:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(177)
		p.Match(BigDuckParserT__5)
	}

	return localctx
}

// INextTypesContext is an interface to support dynamic dispatch.
type INextTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextTypesContext differentiates from other interfaces.
	IsNextTypesContext()
}

type NextTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextTypesContext() *NextTypesContext {
	var p = new(NextTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextTypes
	return p
}

func (*NextTypesContext) IsNextTypesContext() {}

func NewNextTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextTypesContext {
	var p = new(NextTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextTypes

	return p
}

func (s *NextTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *NextTypesContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *NextTypesContext) NextArg() INextArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextArgContext)
}

func (s *NextTypesContext) Var_type() IVar_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *NextTypesContext) NextTypes() INextTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextTypesContext)
}

func (s *NextTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextTypes(s)
	}
}

func (s *NextTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextTypes(s)
	}
}

func (p *BigDuckParser) NextTypes() (localctx INextTypesContext) {
	localctx = NewNextTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BigDuckParserRULE_nextTypes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Match(BigDuckParserT__0)
		}
		{
			p.SetState(180)
			p.Match(BigDuckParserID)
		}
		{
			p.SetState(181)
			p.NextArg()
		}
		{
			p.SetState(182)
			p.Var_type()
		}
		{
			p.SetState(183)
			p.NextTypes()
		}

	case BigDuckParserT__5:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INextArgContext is an interface to support dynamic dispatch.
type INextArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextArgContext differentiates from other interfaces.
	IsNextArgContext()
}

type NextArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextArgContext() *NextArgContext {
	var p = new(NextArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextArg
	return p
}

func (*NextArgContext) IsNextArgContext() {}

func NewNextArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextArgContext {
	var p = new(NextArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextArg

	return p
}

func (s *NextArgContext) GetParser() antlr.Parser { return s.parser }

func (s *NextArgContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *NextArgContext) NextArg() INextArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextArgContext)
}

func (s *NextArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextArg(s)
	}
}

func (s *NextArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextArg(s)
	}
}

func (p *BigDuckParser) NextArg() (localctx INextArgContext) {
	localctx = NewNextArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BigDuckParserRULE_nextArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(192)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(BigDuckParserT__1)
		}
		{
			p.SetState(189)
			p.Match(BigDuckParserID)
		}
		{
			p.SetState(190)
			p.NextArg()
		}

	case BigDuckParserT__2, BigDuckParserINT, BigDuckParserFLOAT, BigDuckParserBOOL:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRet_typeContext is an interface to support dynamic dispatch.
type IRet_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRet_typeContext differentiates from other interfaces.
	IsRet_typeContext()
}

type Ret_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRet_typeContext() *Ret_typeContext {
	var p = new(Ret_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_ret_type
	return p
}

func (*Ret_typeContext) IsRet_typeContext() {}

func NewRet_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ret_typeContext {
	var p = new(Ret_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_ret_type

	return p
}

func (s *Ret_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ret_typeContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *Ret_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ret_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ret_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterRet_type(s)
	}
}

func (s *Ret_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitRet_type(s)
	}
}

func (p *BigDuckParser) Ret_type() (localctx IRet_typeContext) {
	localctx = NewRet_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BigDuckParserRULE_ret_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(BigDuckParserT__6)
	}
	{
		p.SetState(195)
		p.Scalar()
	}

	return localctx
}

// IBool_exprContext is an interface to support dynamic dispatch.
type IBool_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_exprContext differentiates from other interfaces.
	IsBool_exprContext()
}

type Bool_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_exprContext() *Bool_exprContext {
	var p = new(Bool_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_bool_expr
	return p
}

func (*Bool_exprContext) IsBool_exprContext() {}

func NewBool_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_exprContext {
	var p = new(Bool_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_bool_expr

	return p
}

func (s *Bool_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_exprContext) And_expr() IAnd_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_exprContext)
}

func (s *Bool_exprContext) NextBool() INextBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextBoolContext)
}

func (s *Bool_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterBool_expr(s)
	}
}

func (s *Bool_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitBool_expr(s)
	}
}

func (p *BigDuckParser) Bool_expr() (localctx IBool_exprContext) {
	localctx = NewBool_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BigDuckParserRULE_bool_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.And_expr()
	}
	{
		p.SetState(198)
		p.NextBool()
	}

	return localctx
}

// INextBoolContext is an interface to support dynamic dispatch.
type INextBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextBoolContext differentiates from other interfaces.
	IsNextBoolContext()
}

type NextBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextBoolContext() *NextBoolContext {
	var p = new(NextBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextBool
	return p
}

func (*NextBoolContext) IsNextBoolContext() {}

func NewNextBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextBoolContext {
	var p = new(NextBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextBool

	return p
}

func (s *NextBoolContext) GetParser() antlr.Parser { return s.parser }

func (s *NextBoolContext) OR() antlr.TerminalNode {
	return s.GetToken(BigDuckParserOR, 0)
}

func (s *NextBoolContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *NextBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextBool(s)
	}
}

func (s *NextBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextBool(s)
	}
}

func (p *BigDuckParser) NextBool() (localctx INextBoolContext) {
	localctx = NewNextBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BigDuckParserRULE_nextBool)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Match(BigDuckParserOR)
		}
		{
			p.SetState(201)
			p.Bool_expr()
		}

	case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__5, BigDuckParserT__17:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAnd_exprContext is an interface to support dynamic dispatch.
type IAnd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_exprContext differentiates from other interfaces.
	IsAnd_exprContext()
}

type And_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_exprContext() *And_exprContext {
	var p = new(And_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_and_expr
	return p
}

func (*And_exprContext) IsAnd_exprContext() {}

func NewAnd_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_exprContext {
	var p = new(And_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_and_expr

	return p
}

func (s *And_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *And_exprContext) Not_expr() INot_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_exprContext)
}

func (s *And_exprContext) NextAnd() INextAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextAndContext)
}

func (s *And_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterAnd_expr(s)
	}
}

func (s *And_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitAnd_expr(s)
	}
}

func (p *BigDuckParser) And_expr() (localctx IAnd_exprContext) {
	localctx = NewAnd_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BigDuckParserRULE_and_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Not_expr()
	}
	{
		p.SetState(206)
		p.NextAnd()
	}

	return localctx
}

// INextAndContext is an interface to support dynamic dispatch.
type INextAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextAndContext differentiates from other interfaces.
	IsNextAndContext()
}

type NextAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextAndContext() *NextAndContext {
	var p = new(NextAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextAnd
	return p
}

func (*NextAndContext) IsNextAndContext() {}

func NewNextAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextAndContext {
	var p = new(NextAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextAnd

	return p
}

func (s *NextAndContext) GetParser() antlr.Parser { return s.parser }

func (s *NextAndContext) AND() antlr.TerminalNode {
	return s.GetToken(BigDuckParserAND, 0)
}

func (s *NextAndContext) And_expr() IAnd_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_exprContext)
}

func (s *NextAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextAnd(s)
	}
}

func (s *NextAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextAnd(s)
	}
}

func (p *BigDuckParser) NextAnd() (localctx INextAndContext) {
	localctx = NewNextAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BigDuckParserRULE_nextAnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(BigDuckParserAND)
		}
		{
			p.SetState(209)
			p.And_expr()
		}

	case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__5, BigDuckParserT__17, BigDuckParserOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INot_exprContext is an interface to support dynamic dispatch.
type INot_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_exprContext differentiates from other interfaces.
	IsNot_exprContext()
}

type Not_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_exprContext() *Not_exprContext {
	var p = new(Not_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_not_expr
	return p
}

func (*Not_exprContext) IsNot_exprContext() {}

func NewNot_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_exprContext {
	var p = new(Not_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_not_expr

	return p
}

func (s *Not_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Not_exprContext) Bool_term() IBool_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_termContext)
}

func (s *Not_exprContext) NOT() antlr.TerminalNode {
	return s.GetToken(BigDuckParserNOT, 0)
}

func (s *Not_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNot_expr(s)
	}
}

func (s *Not_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNot_expr(s)
	}
}

func (p *BigDuckParser) Not_expr() (localctx INot_exprContext) {
	localctx = NewNot_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BigDuckParserRULE_not_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserNOT:
		{
			p.SetState(213)
			p.Match(BigDuckParserNOT)
		}

	case BigDuckParserT__4, BigDuckParserTRUE, BigDuckParserFALSE, BigDuckParserCTE_INT, BigDuckParserCTE_FLOAT, BigDuckParserID:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(217)
		p.Bool_term()
	}

	return localctx
}

// IBool_termContext is an interface to support dynamic dispatch.
type IBool_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_termContext differentiates from other interfaces.
	IsBool_termContext()
}

type Bool_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_termContext() *Bool_termContext {
	var p = new(Bool_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_bool_term
	return p
}

func (*Bool_termContext) IsBool_termContext() {}

func NewBool_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_termContext {
	var p = new(Bool_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_bool_term

	return p
}

func (s *Bool_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_termContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *Bool_termContext) Rel_expr() IRel_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRel_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRel_exprContext)
}

func (s *Bool_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BigDuckParserTRUE, 0)
}

func (s *Bool_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BigDuckParserFALSE, 0)
}

func (s *Bool_termContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *Bool_termContext) Dim() IDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimContext)
}

func (s *Bool_termContext) Proc_call() IProc_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProc_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProc_callContext)
}

func (s *Bool_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterBool_term(s)
	}
}

func (s *Bool_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitBool_term(s)
	}
}

func (p *BigDuckParser) Bool_term() (localctx IBool_termContext) {
	localctx = NewBool_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BigDuckParserRULE_bool_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Match(BigDuckParserT__4)
		}
		{
			p.SetState(220)
			p.Bool_expr()
		}
		{
			p.SetState(221)
			p.Match(BigDuckParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Rel_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.Match(BigDuckParserTRUE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(225)
			p.Match(BigDuckParserFALSE)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(226)
			p.Match(BigDuckParserID)
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case BigDuckParserT__2:
			{
				p.SetState(227)
				p.Dim()
			}

		case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__5, BigDuckParserT__17, BigDuckParserAND, BigDuckParserOR:

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			p.Proc_call()
		}

	}

	return localctx
}

// IRel_exprContext is an interface to support dynamic dispatch.
type IRel_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRel_exprContext differentiates from other interfaces.
	IsRel_exprContext()
}

type Rel_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRel_exprContext() *Rel_exprContext {
	var p = new(Rel_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_rel_expr
	return p
}

func (*Rel_exprContext) IsRel_exprContext() {}

func NewRel_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rel_exprContext {
	var p = new(Rel_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_rel_expr

	return p
}

func (s *Rel_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Rel_exprContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *Rel_exprContext) OpRel() IOpRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpRelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpRelContext)
}

func (s *Rel_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rel_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rel_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterRel_expr(s)
	}
}

func (s *Rel_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitRel_expr(s)
	}
}

func (p *BigDuckParser) Rel_expr() (localctx IRel_exprContext) {
	localctx = NewRel_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BigDuckParserRULE_rel_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Num_expr()
	}
	{
		p.SetState(235)
		p.OpRel()
	}

	return localctx
}

// IOpRelContext is an interface to support dynamic dispatch.
type IOpRelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpRelContext differentiates from other interfaces.
	IsOpRelContext()
}

type OpRelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpRelContext() *OpRelContext {
	var p = new(OpRelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_opRel
	return p
}

func (*OpRelContext) IsOpRelContext() {}

func NewOpRelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpRelContext {
	var p = new(OpRelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_opRel

	return p
}

func (s *OpRelContext) GetParser() antlr.Parser { return s.parser }

func (s *OpRelContext) RelOp() IRelOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelOpContext)
}

func (s *OpRelContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *OpRelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpRelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpRelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterOpRel(s)
	}
}

func (s *OpRelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitOpRel(s)
	}
}

func (p *BigDuckParser) OpRel() (localctx IOpRelContext) {
	localctx = NewOpRelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BigDuckParserRULE_opRel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.RelOp()
	}
	{
		p.SetState(238)
		p.Num_expr()
	}

	return localctx
}

// IRelOpContext is an interface to support dynamic dispatch.
type IRelOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelOpContext differentiates from other interfaces.
	IsRelOpContext()
}

type RelOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelOpContext() *RelOpContext {
	var p = new(RelOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_relOp
	return p
}

func (*RelOpContext) IsRelOpContext() {}

func NewRelOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelOpContext {
	var p = new(RelOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_relOp

	return p
}

func (s *RelOpContext) GetParser() antlr.Parser { return s.parser }
func (s *RelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterRelOp(s)
	}
}

func (s *RelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitRelOp(s)
	}
}

func (p *BigDuckParser) RelOp() (localctx IRelOpContext) {
	localctx = NewRelOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BigDuckParserRULE_relOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BigDuckParserT__7)|(1<<BigDuckParserT__8)|(1<<BigDuckParserT__9)|(1<<BigDuckParserT__10)|(1<<BigDuckParserT__11)|(1<<BigDuckParserT__12))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INum_exprContext is an interface to support dynamic dispatch.
type INum_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNum_exprContext differentiates from other interfaces.
	IsNum_exprContext()
}

type Num_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNum_exprContext() *Num_exprContext {
	var p = new(Num_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_num_expr
	return p
}

func (*Num_exprContext) IsNum_exprContext() {}

func NewNum_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Num_exprContext {
	var p = new(Num_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_num_expr

	return p
}

func (s *Num_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Num_exprContext) Prod_expr() IProd_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProd_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProd_exprContext)
}

func (s *Num_exprContext) NextSum() INextSumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextSumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextSumContext)
}

func (s *Num_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Num_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Num_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNum_expr(s)
	}
}

func (s *Num_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNum_expr(s)
	}
}

func (p *BigDuckParser) Num_expr() (localctx INum_exprContext) {
	localctx = NewNum_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BigDuckParserRULE_num_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Prod_expr()
	}
	{
		p.SetState(243)
		p.NextSum()
	}

	return localctx
}

// INextSumContext is an interface to support dynamic dispatch.
type INextSumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextSumContext differentiates from other interfaces.
	IsNextSumContext()
}

type NextSumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextSumContext() *NextSumContext {
	var p = new(NextSumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextSum
	return p
}

func (*NextSumContext) IsNextSumContext() {}

func NewNextSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextSumContext {
	var p = new(NextSumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextSum

	return p
}

func (s *NextSumContext) GetParser() antlr.Parser { return s.parser }

func (s *NextSumContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *NextSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextSumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextSum(s)
	}
}

func (s *NextSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextSum(s)
	}
}

func (p *BigDuckParser) NextSum() (localctx INextSumContext) {
	localctx = NewNextSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BigDuckParserRULE_nextSum)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__13, BigDuckParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BigDuckParserT__13 || _la == BigDuckParserT__14) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(246)
			p.Num_expr()
		}

	case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__3, BigDuckParserT__5, BigDuckParserT__7, BigDuckParserT__8, BigDuckParserT__9, BigDuckParserT__10, BigDuckParserT__11, BigDuckParserT__12, BigDuckParserT__17, BigDuckParserAND, BigDuckParserOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProd_exprContext is an interface to support dynamic dispatch.
type IProd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProd_exprContext differentiates from other interfaces.
	IsProd_exprContext()
}

type Prod_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProd_exprContext() *Prod_exprContext {
	var p = new(Prod_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_prod_expr
	return p
}

func (*Prod_exprContext) IsProd_exprContext() {}

func NewProd_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prod_exprContext {
	var p = new(Prod_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_prod_expr

	return p
}

func (s *Prod_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Prod_exprContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *Prod_exprContext) NextProd() INextProdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextProdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextProdContext)
}

func (s *Prod_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prod_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prod_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterProd_expr(s)
	}
}

func (s *Prod_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitProd_expr(s)
	}
}

func (p *BigDuckParser) Prod_expr() (localctx IProd_exprContext) {
	localctx = NewProd_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BigDuckParserRULE_prod_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Factor()
	}
	{
		p.SetState(251)
		p.NextProd()
	}

	return localctx
}

// INextProdContext is an interface to support dynamic dispatch.
type INextProdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextProdContext differentiates from other interfaces.
	IsNextProdContext()
}

type NextProdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextProdContext() *NextProdContext {
	var p = new(NextProdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextProd
	return p
}

func (*NextProdContext) IsNextProdContext() {}

func NewNextProdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextProdContext {
	var p = new(NextProdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextProd

	return p
}

func (s *NextProdContext) GetParser() antlr.Parser { return s.parser }

func (s *NextProdContext) Prod_expr() IProd_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProd_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProd_exprContext)
}

func (s *NextProdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextProdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextProdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextProd(s)
	}
}

func (s *NextProdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextProd(s)
	}
}

func (p *BigDuckParser) NextProd() (localctx INextProdContext) {
	localctx = NewNextProdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BigDuckParserRULE_nextProd)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__15, BigDuckParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BigDuckParserT__15 || _la == BigDuckParserT__16) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(254)
			p.Prod_expr()
		}

	case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__3, BigDuckParserT__5, BigDuckParserT__7, BigDuckParserT__8, BigDuckParserT__9, BigDuckParserT__10, BigDuckParserT__11, BigDuckParserT__12, BigDuckParserT__13, BigDuckParserT__14, BigDuckParserT__17, BigDuckParserAND, BigDuckParserOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *FactorContext) Dim() IDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimContext)
}

func (s *FactorContext) CTE_INT() antlr.TerminalNode {
	return s.GetToken(BigDuckParserCTE_INT, 0)
}

func (s *FactorContext) CTE_FLOAT() antlr.TerminalNode {
	return s.GetToken(BigDuckParserCTE_FLOAT, 0)
}

func (s *FactorContext) Proc_call() IProc_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProc_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProc_callContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *BigDuckParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BigDuckParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(BigDuckParserT__4)
		}
		{
			p.SetState(259)
			p.Num_expr()
		}
		{
			p.SetState(260)
			p.Match(BigDuckParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(BigDuckParserID)
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case BigDuckParserT__2:
			{
				p.SetState(263)
				p.Dim()
			}

		case BigDuckParserT__0, BigDuckParserT__1, BigDuckParserT__3, BigDuckParserT__5, BigDuckParserT__7, BigDuckParserT__8, BigDuckParserT__9, BigDuckParserT__10, BigDuckParserT__11, BigDuckParserT__12, BigDuckParserT__13, BigDuckParserT__14, BigDuckParserT__15, BigDuckParserT__16, BigDuckParserT__17, BigDuckParserAND, BigDuckParserOR:

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(267)
			p.Match(BigDuckParserCTE_INT)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(268)
			p.Match(BigDuckParserCTE_FLOAT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(269)
			p.Proc_call()
		}

	}

	return localctx
}

// IProc_callContext is an interface to support dynamic dispatch.
type IProc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProc_callContext differentiates from other interfaces.
	IsProc_callContext()
}

type Proc_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_callContext() *Proc_callContext {
	var p = new(Proc_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_proc_call
	return p
}

func (*Proc_callContext) IsProc_callContext() {}

func NewProc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_callContext {
	var p = new(Proc_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_proc_call

	return p
}

func (s *Proc_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_callContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *Proc_callContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Proc_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterProc_call(s)
	}
}

func (s *Proc_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitProc_call(s)
	}
}

func (p *BigDuckParser) Proc_call() (localctx IProc_callContext) {
	localctx = NewProc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BigDuckParserRULE_proc_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(BigDuckParserID)
	}
	{
		p.SetState(273)
		p.Match(BigDuckParserT__4)
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__4, BigDuckParserNOT, BigDuckParserTRUE, BigDuckParserFALSE, BigDuckParserCTE_INT, BigDuckParserCTE_FLOAT, BigDuckParserCTE_STRING, BigDuckParserID:
		{
			p.SetState(274)
			p.Param()
		}

	case BigDuckParserT__5:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(278)
		p.Match(BigDuckParserT__5)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *ParamContext) NextParam() INextParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextParamContext)
}

func (s *ParamContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *ParamContext) CTE_STRING() antlr.TerminalNode {
	return s.GetToken(BigDuckParserCTE_STRING, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *BigDuckParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BigDuckParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Bool_expr()
		}
		{
			p.SetState(281)
			p.NextParam()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Num_expr()
		}
		{
			p.SetState(284)
			p.NextParam()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(286)
			p.Match(BigDuckParserCTE_STRING)
		}
		{
			p.SetState(287)
			p.NextParam()
		}

	}

	return localctx
}

// INextParamContext is an interface to support dynamic dispatch.
type INextParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextParamContext differentiates from other interfaces.
	IsNextParamContext()
}

type NextParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextParamContext() *NextParamContext {
	var p = new(NextParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_nextParam
	return p
}

func (*NextParamContext) IsNextParamContext() {}

func NewNextParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextParamContext {
	var p = new(NextParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_nextParam

	return p
}

func (s *NextParamContext) GetParser() antlr.Parser { return s.parser }

func (s *NextParamContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *NextParamContext) NextParam() INextParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextParamContext)
}

func (s *NextParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterNextParam(s)
	}
}

func (s *NextParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitNextParam(s)
	}
}

func (p *BigDuckParser) NextParam() (localctx INextParamContext) {
	localctx = NewNextParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BigDuckParserRULE_nextParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Match(BigDuckParserT__1)
		}
		{
			p.SetState(291)
			p.Param()
		}
		{
			p.SetState(292)
			p.NextParam()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BigDuckParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BigDuckParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(BigDuckParserT__17)
	}
	{
		p.SetState(298)
		p.Stmts()
	}
	{
		p.SetState(299)
		p.Match(BigDuckParserT__18)
	}

	return localctx
}

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtsContext differentiates from other interfaces.
	IsStmtsContext()
}

type StmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtsContext() *StmtsContext {
	var p = new(StmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_stmts
	return p
}

func (*StmtsContext) IsStmtsContext() {}

func NewStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtsContext {
	var p = new(StmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_stmts

	return p
}

func (s *StmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtsContext) Stmt() IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtsContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterStmts(s)
	}
}

func (s *StmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitStmts(s)
	}
}

func (p *BigDuckParser) Stmts() (localctx IStmtsContext) {
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BigDuckParserRULE_stmts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserRETURN, BigDuckParserIF, BigDuckParserLOOP, BigDuckParserBREAK, BigDuckParserSKIP_W, BigDuckParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Stmt()
		}
		{
			p.SetState(302)
			p.Stmts()
		}

	case BigDuckParserT__18:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *StmtContext) Loop_stmt() ILoop_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoop_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoop_stmtContext)
}

func (s *StmtContext) Ctrl_flow() ICtrl_flowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICtrl_flowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICtrl_flowContext)
}

func (s *StmtContext) Ret_stmt() IRet_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRet_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRet_stmtContext)
}

func (s *StmtContext) Proc_call() IProc_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProc_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProc_callContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *BigDuckParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BigDuckParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Assignment()
		}
		{
			p.SetState(308)
			p.Match(BigDuckParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Condition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Loop_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Ctrl_flow()
		}
		{
			p.SetState(313)
			p.Match(BigDuckParserT__0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(315)
			p.Ret_stmt()
		}
		{
			p.SetState(316)
			p.Match(BigDuckParserT__0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(318)
			p.Proc_call()
		}
		{
			p.SetState(319)
			p.Match(BigDuckParserT__0)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(BigDuckParserID, 0)
}

func (s *AssignmentContext) Dim() IDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimContext)
}

func (s *AssignmentContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *AssignmentContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *BigDuckParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BigDuckParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(BigDuckParserID)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserT__2:
		{
			p.SetState(324)
			p.Dim()
		}

	case BigDuckParserT__19:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(328)
		p.Match(BigDuckParserT__19)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(329)
			p.Num_expr()
		}

	case 2:
		{
			p.SetState(330)
			p.Bool_expr()
		}

	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) IF() antlr.TerminalNode {
	return s.GetToken(BigDuckParserIF, 0)
}

func (s *ConditionContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *ConditionContext) BodyCond() IBodyCondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyCondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyCondContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *BigDuckParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BigDuckParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(BigDuckParserIF)
	}
	{
		p.SetState(334)
		p.Bool_expr()
	}
	{
		p.SetState(335)
		p.BodyCond()
	}

	return localctx
}

// IBodyCondContext is an interface to support dynamic dispatch.
type IBodyCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyCondContext differentiates from other interfaces.
	IsBodyCondContext()
}

type BodyCondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyCondContext() *BodyCondContext {
	var p = new(BodyCondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_bodyCond
	return p
}

func (*BodyCondContext) IsBodyCondContext() {}

func NewBodyCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyCondContext {
	var p = new(BodyCondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_bodyCond

	return p
}

func (s *BodyCondContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyCondContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BodyCondContext) EndIfBlock() IEndIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndIfBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndIfBlockContext)
}

func (s *BodyCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterBodyCond(s)
	}
}

func (s *BodyCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitBodyCond(s)
	}
}

func (p *BigDuckParser) BodyCond() (localctx IBodyCondContext) {
	localctx = NewBodyCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BigDuckParserRULE_bodyCond)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Block()
	}
	{
		p.SetState(338)
		p.EndIfBlock()
	}

	return localctx
}

// IEndIfBlockContext is an interface to support dynamic dispatch.
type IEndIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndIfBlockContext differentiates from other interfaces.
	IsEndIfBlockContext()
}

type EndIfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndIfBlockContext() *EndIfBlockContext {
	var p = new(EndIfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_endIfBlock
	return p
}

func (*EndIfBlockContext) IsEndIfBlockContext() {}

func NewEndIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndIfBlockContext {
	var p = new(EndIfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_endIfBlock

	return p
}

func (s *EndIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *EndIfBlockContext) Alter() IAlterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterContext)
}

func (s *EndIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterEndIfBlock(s)
	}
}

func (s *EndIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitEndIfBlock(s)
	}
}

func (p *BigDuckParser) EndIfBlock() (localctx IEndIfBlockContext) {
	localctx = NewEndIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BigDuckParserRULE_endIfBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserELSE:
		{
			p.SetState(340)
			p.Alter()
		}

	case BigDuckParserT__18, BigDuckParserRETURN, BigDuckParserIF, BigDuckParserLOOP, BigDuckParserBREAK, BigDuckParserSKIP_W, BigDuckParserID:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAlterContext is an interface to support dynamic dispatch.
type IAlterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterContext differentiates from other interfaces.
	IsAlterContext()
}

type AlterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterContext() *AlterContext {
	var p = new(AlterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_alter
	return p
}

func (*AlterContext) IsAlterContext() {}

func NewAlterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterContext {
	var p = new(AlterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_alter

	return p
}

func (s *AlterContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterContext) ELSE() antlr.TerminalNode {
	return s.GetToken(BigDuckParserELSE, 0)
}

func (s *AlterContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *AlterContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *AlterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterAlter(s)
	}
}

func (s *AlterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitAlter(s)
	}
}

func (p *BigDuckParser) Alter() (localctx IAlterContext) {
	localctx = NewAlterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BigDuckParserRULE_alter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(BigDuckParserELSE)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserIF:
		{
			p.SetState(345)
			p.Condition()
		}

	case BigDuckParserT__17:
		{
			p.SetState(346)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoop_stmtContext is an interface to support dynamic dispatch.
type ILoop_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoop_stmtContext differentiates from other interfaces.
	IsLoop_stmtContext()
}

type Loop_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_stmtContext() *Loop_stmtContext {
	var p = new(Loop_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_loop_stmt
	return p
}

func (*Loop_stmtContext) IsLoop_stmtContext() {}

func NewLoop_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_stmtContext {
	var p = new(Loop_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_loop_stmt

	return p
}

func (s *Loop_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_stmtContext) LOOP() antlr.TerminalNode {
	return s.GetToken(BigDuckParserLOOP, 0)
}

func (s *Loop_stmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Loop_stmtContext) ForStyle() IForStyleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStyleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStyleContext)
}

func (s *Loop_stmtContext) WhileStyle() IWhileStyleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStyleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStyleContext)
}

func (s *Loop_stmtContext) InfLoop() IInfLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfLoopContext)
}

func (s *Loop_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterLoop_stmt(s)
	}
}

func (s *Loop_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitLoop_stmt(s)
	}
}

func (p *BigDuckParser) Loop_stmt() (localctx ILoop_stmtContext) {
	localctx = NewLoop_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BigDuckParserRULE_loop_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(BigDuckParserLOOP)
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(350)
			p.ForStyle()
		}

	case 2:
		{
			p.SetState(351)
			p.WhileStyle()
		}

	case 3:
		{
			p.SetState(352)
			p.InfLoop()
		}

	}
	{
		p.SetState(355)
		p.Block()
	}

	return localctx
}

// IForStyleContext is an interface to support dynamic dispatch.
type IForStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStyleContext differentiates from other interfaces.
	IsForStyleContext()
}

type ForStyleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStyleContext() *ForStyleContext {
	var p = new(ForStyleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_forStyle
	return p
}

func (*ForStyleContext) IsForStyleContext() {}

func NewForStyleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStyleContext {
	var p = new(ForStyleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_forStyle

	return p
}

func (s *ForStyleContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStyleContext) ForCond() IForCondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForCondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForCondContext)
}

func (s *ForStyleContext) Ctrl_var() ICtrl_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICtrl_varContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICtrl_varContext)
}

func (s *ForStyleContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForStyleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStyleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStyleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterForStyle(s)
	}
}

func (s *ForStyleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitForStyle(s)
	}
}

func (p *BigDuckParser) ForStyle() (localctx IForStyleContext) {
	localctx = NewForStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BigDuckParserRULE_forStyle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BigDuckParserID:
		{
			p.SetState(357)
			p.Assignment()
		}

	case BigDuckParserT__0:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(361)
		p.Match(BigDuckParserT__0)
	}
	{
		p.SetState(362)
		p.ForCond()
	}
	{
		p.SetState(363)
		p.Match(BigDuckParserT__0)
	}
	{
		p.SetState(364)
		p.Ctrl_var()
	}

	return localctx
}

// IForCondContext is an interface to support dynamic dispatch.
type IForCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForCondContext differentiates from other interfaces.
	IsForCondContext()
}

type ForCondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForCondContext() *ForCondContext {
	var p = new(ForCondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_forCond
	return p
}

func (*ForCondContext) IsForCondContext() {}

func NewForCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForCondContext {
	var p = new(ForCondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_forCond

	return p
}

func (s *ForCondContext) GetParser() antlr.Parser { return s.parser }

func (s *ForCondContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *ForCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterForCond(s)
	}
}

func (s *ForCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitForCond(s)
	}
}

func (p *BigDuckParser) ForCond() (localctx IForCondContext) {
	localctx = NewForCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BigDuckParserRULE_forCond)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Bool_expr()
	}

	return localctx
}

// ICtrl_varContext is an interface to support dynamic dispatch.
type ICtrl_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCtrl_varContext differentiates from other interfaces.
	IsCtrl_varContext()
}

type Ctrl_varContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtrl_varContext() *Ctrl_varContext {
	var p = new(Ctrl_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_ctrl_var
	return p
}

func (*Ctrl_varContext) IsCtrl_varContext() {}

func NewCtrl_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctrl_varContext {
	var p = new(Ctrl_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_ctrl_var

	return p
}

func (s *Ctrl_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctrl_varContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Ctrl_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctrl_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctrl_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterCtrl_var(s)
	}
}

func (s *Ctrl_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitCtrl_var(s)
	}
}

func (p *BigDuckParser) Ctrl_var() (localctx ICtrl_varContext) {
	localctx = NewCtrl_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BigDuckParserRULE_ctrl_var)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Assignment()
	}

	return localctx
}

// IWhileStyleContext is an interface to support dynamic dispatch.
type IWhileStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStyleContext differentiates from other interfaces.
	IsWhileStyleContext()
}

type WhileStyleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStyleContext() *WhileStyleContext {
	var p = new(WhileStyleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_whileStyle
	return p
}

func (*WhileStyleContext) IsWhileStyleContext() {}

func NewWhileStyleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStyleContext {
	var p = new(WhileStyleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_whileStyle

	return p
}

func (s *WhileStyleContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStyleContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *WhileStyleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStyleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStyleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterWhileStyle(s)
	}
}

func (s *WhileStyleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitWhileStyle(s)
	}
}

func (p *BigDuckParser) WhileStyle() (localctx IWhileStyleContext) {
	localctx = NewWhileStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BigDuckParserRULE_whileStyle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Bool_expr()
	}

	return localctx
}

// IInfLoopContext is an interface to support dynamic dispatch.
type IInfLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfLoopContext differentiates from other interfaces.
	IsInfLoopContext()
}

type InfLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfLoopContext() *InfLoopContext {
	var p = new(InfLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_infLoop
	return p
}

func (*InfLoopContext) IsInfLoopContext() {}

func NewInfLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfLoopContext {
	var p = new(InfLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_infLoop

	return p
}

func (s *InfLoopContext) GetParser() antlr.Parser { return s.parser }
func (s *InfLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterInfLoop(s)
	}
}

func (s *InfLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitInfLoop(s)
	}
}

func (p *BigDuckParser) InfLoop() (localctx IInfLoopContext) {
	localctx = NewInfLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BigDuckParserRULE_infLoop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)

	return localctx
}

// ICtrl_flowContext is an interface to support dynamic dispatch.
type ICtrl_flowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCtrl_flowContext differentiates from other interfaces.
	IsCtrl_flowContext()
}

type Ctrl_flowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtrl_flowContext() *Ctrl_flowContext {
	var p = new(Ctrl_flowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_ctrl_flow
	return p
}

func (*Ctrl_flowContext) IsCtrl_flowContext() {}

func NewCtrl_flowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctrl_flowContext {
	var p = new(Ctrl_flowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_ctrl_flow

	return p
}

func (s *Ctrl_flowContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctrl_flowContext) BREAK() antlr.TerminalNode {
	return s.GetToken(BigDuckParserBREAK, 0)
}

func (s *Ctrl_flowContext) SKIP_W() antlr.TerminalNode {
	return s.GetToken(BigDuckParserSKIP_W, 0)
}

func (s *Ctrl_flowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctrl_flowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctrl_flowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterCtrl_flow(s)
	}
}

func (s *Ctrl_flowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitCtrl_flow(s)
	}
}

func (p *BigDuckParser) Ctrl_flow() (localctx ICtrl_flowContext) {
	localctx = NewCtrl_flowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BigDuckParserRULE_ctrl_flow)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BigDuckParserBREAK || _la == BigDuckParserSKIP_W) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRet_stmtContext is an interface to support dynamic dispatch.
type IRet_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRet_stmtContext differentiates from other interfaces.
	IsRet_stmtContext()
}

type Ret_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRet_stmtContext() *Ret_stmtContext {
	var p = new(Ret_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BigDuckParserRULE_ret_stmt
	return p
}

func (*Ret_stmtContext) IsRet_stmtContext() {}

func NewRet_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ret_stmtContext {
	var p = new(Ret_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigDuckParserRULE_ret_stmt

	return p
}

func (s *Ret_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Ret_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BigDuckParserRETURN, 0)
}

func (s *Ret_stmtContext) Num_expr() INum_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_exprContext)
}

func (s *Ret_stmtContext) Bool_expr() IBool_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *Ret_stmtContext) Proc_call() IProc_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProc_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProc_callContext)
}

func (s *Ret_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ret_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ret_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.EnterRet_stmt(s)
	}
}

func (s *Ret_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigDuckListener); ok {
		listenerT.ExitRet_stmt(s)
	}
}

func (p *BigDuckParser) Ret_stmt() (localctx IRet_stmtContext) {
	localctx = NewRet_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BigDuckParserRULE_ret_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(BigDuckParserRETURN)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(377)
			p.Num_expr()
		}

	case 2:
		{
			p.SetState(378)
			p.Bool_expr()
		}

	case 3:
		{
			p.SetState(379)
			p.Proc_call()
		}

	case 4:

	}

	return localctx
}
