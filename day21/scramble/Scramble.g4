grammar Scramble;

scrambler
  : operation (Newline operation)* Newline?
  ;

operation
  : swapPosition
  | swapLetter
  | rotateLeft
  | rotateRight
  | rotateLetter
  | reverse
  | move
  ;

swapPosition
  : 'swap position ' Num ' with position ' Num
  ;

swapLetter
  : 'swap letter ' Letter ' with letter ' Letter
  ;

rotateLeft
  : 'rotate left ' Num ( ' steps' | ' step' )
  ;

rotateRight
  : 'rotate right ' Num ( ' steps' | ' step' )
  ;

rotateLetter
  : 'rotate based on position of letter ' Letter
  ;

reverse
  : 'reverse positions ' Num ' through ' Num
  ;

move
  : 'move position ' Num ' to position ' Num
  ;

Num : [0-9]+ ;

Letter : [a-z]+ ;

Newline : [\r\n]+ ;
