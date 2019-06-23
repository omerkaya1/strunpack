# strunpack - Single-method package that unpack a provided string in the following manner:

StringParse("a4bc2d5e") => "aaaabccddddde"
StringParse("abcd")     => "abcd"
StringParse("45")       => "" // strings that contain only numeric characters are not valid
StringParse(`qwe\4\5`)  => `qwe45`
StringParse(`qwe\45`)   => `qwe44444`
StringParse(`qwe\\5`)   => `qwe\\\\\`
