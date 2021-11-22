" Vim syntax file
" Language:         BigDuck
" Maintainer:       Jair Antonio
" Latest Revision:  10 November 2021

if exists("b:current_syntax")
    finish
endif

syn keyword duckKeywords    proc    return  if      else
syn keyword duckKeywords    loop    break   skip    and
syn keyword duckKeywords    or      not     var     int
syn keyword duckKeywords    float   bool    true    false

syn keyword duckBuiltin     print   read
syn keyword duckBuiltin     sin     asin    cos     acos    tan     atan
syn keyword duckBuiltin     atan2
syn keyword duckBuiltin     exp     ln      sqrt    pow     log     mod
syn keyword duckBuiltin     abs     ceil    floor   mean    median  mode

syn match duckNumber '\<\d\+'
syn match duckNumber '\<\d\+\.\d*\([eE][-+]\=\d\+\)\=[fl]\=' 
syn match duckNumber '\<\d\+[eE][-+]\=\d\+[fl]\=\>'

syn region duckString start='"' end='"' 
syn region duckComment start='#|' end='|#' 

let b:current_syntax = "duck"

hi def link duckKeywords    Statement
hi def link duckComment     Comment
hi def link duckNumber      Number
hi def link duckString      String
hi def link duckBuiltin     Function
