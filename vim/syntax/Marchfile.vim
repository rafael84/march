syntax keyword marchKeywords config service domain default
syntax keyword marchDataTypes string int bool hash timestamp
syntax match marchComment "#.*$"
syntax region marchString start='"' end='"' contained

hi link marchKeywords Keyword
hi link marchDataTypes Type
hi link marchComment Comment
hi link marchString	String
