GOOF----LE-4-2.0 G      ] � 4        h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  language�	g  glil�	 �		g  filenameS�	
f  language/glil.scm�	g  importsS�	g  system�	g  base�	g  syntax�	 �	 �	g  pmatch�	 �	 �	g  srfi�	g  srfi-1�	 �	g  selectS�	g  fold�	 �	 �	 �	g  exportsS�	g  <glil-program>�	g  make-glil-program�	g  glil-program?�	 g  glil-program-meta�	!g  glil-program-body�	"g  <glil-std-prelude>�	#g  make-glil-std-prelude�	$g  glil-std-prelude?�	%g  glil-std-prelude-nreq�	&g  glil-std-prelude-nlocs�	'g  glil-std-prelude-else-label�	(g  <glil-opt-prelude>�	)g  make-glil-opt-prelude�	*g  glil-opt-prelude?�	+g  glil-opt-prelude-nreq�	,g  glil-opt-prelude-nopt�	-g  glil-opt-prelude-rest�	.g  glil-opt-prelude-nlocs�	/g  glil-opt-prelude-else-label�	0g  <glil-kw-prelude>�	1g  make-glil-kw-prelude�	2g  glil-kw-prelude?�	3g  glil-kw-prelude-nreq�	4g  glil-kw-prelude-nopt�	5g  glil-kw-prelude-kw�	6g  !glil-kw-prelude-allow-other-keys?�	7g  glil-kw-prelude-rest�	8g  glil-kw-prelude-nlocs�	9g  glil-kw-prelude-else-label�	:g  <glil-bind>�	;g  make-glil-bind�	<g  
glil-bind?�	=g  glil-bind-vars�	>g  <glil-mv-bind>�	?g  make-glil-mv-bind�	@g  glil-mv-bind?�	Ag  glil-mv-bind-vars�	Bg  glil-mv-bind-rest�	Cg  <glil-unbind>�	Dg  make-glil-unbind�	Eg  glil-unbind?�	Fg  <glil-source>�	Gg  make-glil-source�	Hg  glil-source?�	Ig  glil-source-props�	Jg  <glil-void>�	Kg  make-glil-void�	Lg  
glil-void?�	Mg  <glil-const>�	Ng  make-glil-const�	Og  glil-const?�	Pg  glil-const-obj�	Qg  <glil-lexical>�	Rg  make-glil-lexical�	Sg  glil-lexical?�	Tg  glil-lexical-local?�	Ug  glil-lexical-boxed?�	Vg  glil-lexical-op�	Wg  glil-lexical-index�	Xg  <glil-toplevel>�	Yg  make-glil-toplevel�	Zg  glil-toplevel?�	[g  glil-toplevel-op�	\g  glil-toplevel-name�	]g  <glil-module>�	^g  make-glil-module�	_g  glil-module?�	`g  glil-module-op�	ag  glil-module-mod�	bg  glil-module-name�	cg  glil-module-public?�	dg  <glil-label>�	eg  make-glil-label�	fg  glil-label?�	gg  glil-label-label�	hg  <glil-branch>�	ig  make-glil-branch�	jg  glil-branch?�	kg  glil-branch-inst�	lg  glil-branch-label�	mg  <glil-call>�	ng  make-glil-call�	og  
glil-call?�	pg  glil-call-inst�	qg  glil-call-nargs�	rg  <glil-mv-call>�	sg  make-glil-mv-call�	tg  glil-mv-call?�	ug  glil-mv-call-nargs�	vg  glil-mv-call-ra�	wg  <glil-prompt>�	xg  make-glil-prompt�	yg  glil-prompt?�	zg  glil-prompt-label�	{g  glil-prompt-escape-only?�	|g  
parse-glil�	}g  unparse-glil�	~ !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|} a�	g  set-current-module� � � � � �g  format� �f  
#<glil ~s>� �g  
print-glil� �g  make-record-type� �f  <glil-program>� �g  meta� �g  body� ��� � �g  record-predicate� �g  make-procedure-with-setter� �g  record-accessor� �g  record-modifier� �f  <glil-std-prelude>� �g  nreq� �g  nlocs� �g  
else-label� ���� � �f  <glil-opt-prelude>� �g  nopt� �g  rest� ������ � �f  <glil-kw-prelude>� �g  kw� �g  allow-other-keys?� �������� � �f  <glil-bind>� �g  vars� �� � �f  <glil-mv-bind>� ��� � �f  <glil-unbind>� �f  <glil-source>� �g  props� �� � �f  <glil-void>� �f  <glil-const>� �g  obj� �� � �f  <glil-lexical>� �g  local?� �g  boxed?� �g  op� �g  index� ����� � �f  <glil-toplevel>� �g  name� ��� � �f  <glil-module>� �g  mod� �g  public?� ����� � �f  <glil-label>� �g  label� �� � �f  <glil-branch>� �g  inst� ��� � �f  <glil-call>� �g  nargs� ��� � �f  <glil-mv-call>� �g  ra� ��� � �f  <glil-prompt>� �g  escape-only?� ��� � �g  prompt� �g  error� �f  invalid glil� �g  mv-call� �g  call� �g  branch� �g  module� �g  private� �g  public� �g  toplevel� �g  lexical� �g  const� �g  void� �g  source� �g  unbind� �g  mv-bind� �g  bind� �g  
kw-prelude� �g  opt-prelude� �g  std-prelude� �g  program� �g  map� �� � �� � �� � �f  unhandled record�C 5  h�3  v   ]4	
~5 4� >  "  G   ��}  h   �   ]4 56       �       g  x
		 g  port		  g  filenamef  language/glil.scm�
	Q
��		R	��			R	��		R	�� 			  g  nameg  
print-glil� C�R4�i���i5R h   �   - 1 3  � C     �       g  meta
			 g  body			 g   g102913				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-program� CR4�ii5R4�i4�ii�54�ii�55 R4�i4�ii�54�ii�55!R4�i���i5"R"        h   �   - 1 3  � C   �       g  nreq
			 g  nlocs			 g  
else-label				 g   g102915				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-std-prelude� C#R4�i"i5$R4�i4�i"i�54�i"i�55%R4�i4�i"i�54�i"i�55&R4�i4�i"i�54�i"i�55'R4�i���i5(R(       h    �   - 1 3  � C       �       g  nreq
			 g  nopt			 g  rest				 g  nlocs				 g  
else-label				 g   g102917				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-opt-prelude� C)R4�i(i5*R4�i4�i(i�54�i(i�55+R4�i4�i(i�54�i(i�55,R4�i4�i(i�54�i(i�55-R4�i4�i(i�54�i(i�55.R4�i4�i(i�54�i(i�55/R4�i���i50R0   h      - 1 3  � C         g  nreq
			 g  nopt			 g  rest				 g  kw				 g  allow-other-keys?				 g  nlocs				 g  
else-label				 g   g102919				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-kw-prelude� C1R4�i0i52R4�i4�i0i�54�i0i�553R4�i4�i0i�54�i0i�554R4�i4�i0i�54�i0i�557R4�i4�i0i�54�i0i�555R4�i4�i0i�54�i0i�556R4�i4�i0i�54�i0i�558R4�i4�i0i�54�i0i�559R4�i���i5:R:     h   �   - 1 3  � C       �       g  vars
			 g   g102921			  g  filenamef  language/glil.scm�
	T
�� 			
  g  nameg  make-glil-bind� C;R4�i:i5<R4�i4�i:i�54�i:i�55=R4�i���i5>R>   h   �   - 1 3  � C     �       g  vars
			 g  rest			 g   g102923				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-mv-bind� C?R4�i>i5@R4�i4�i>i�54�i>i�55AR4�i4�i>i�54�i>i�55BR4�i��i5CRC h   y   -  1  3 � C q       g   g102925
			  g  filenamef  language/glil.scm�
	T
�� 			


  g  nameg  make-glil-unbind� CDR4�iCi5ER4�i���i5FRF      h   �   - 1 3  � C       �       g  props
			 g   g102927			  g  filenamef  language/glil.scm�
	T
�� 			
  g  nameg  make-glil-source� CGR4�iFi5HR4�i4�iFi�54�iFi�55IR4�i��i5JRJ h   w   -  1  3 � C o       g   g102929
			  g  filenamef  language/glil.scm�
	T
�� 			


  g  nameg  make-glil-void� CKR4�iJi5LR4�i���i5MRM        h   �   - 1 3  � C       �       g  obj
			 g   g102931			  g  filenamef  language/glil.scm�
	T
�� 			
  g  nameg  make-glil-const� CNR4�iMi5OR4�i4�iMi�54�iMi�55PR4�i���i5QRQ   h   �   - 1 3  � C �       g  local?
			 g  boxed?			 g  op				 g  index				 g   g102933				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-lexical� CRR4�iQi5SR4�i4�iQi�54�iQi�55TR4�i4�iQi�54�iQi�55UR4�i4�iQi�54�iQi�55VR4�i4�iQi�54�iQi�55WR4�i���i5XRX h   �   - 1 3  � C     �       g  op
			 g  name			 g   g102935				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-toplevel� CYR4�iXi5ZR4�i4�iXi�54�iXi�55[R4�i4�iXi�54�iXi�55\R4�i���i5]R] h   �   - 1 3  � C �       g  op
			 g  mod			 g  name				 g  public?				 g   g102937				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-module� C^R4�i]i5_R4�i4�i]i�54�i]i�55`R4�i4�i]i�54�i]i�55aR4�i4�i]i�54�i]i�55bR4�i4�i]i�54�i]i�55cR4�i���i5dRd     h   �   - 1 3  � C       �       g  label
			 g   g102939			  g  filenamef  language/glil.scm�
	T
�� 			
  g  nameg  make-glil-label� CeR4�idi5fR4�i4�idi�54�idi�55gR4�i���i5hRh h   �   - 1 3  � C     �       g  inst
			 g  label			 g   g102941				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-branch� CiR4�ihi5jR4�i4�ihi�54�ihi�55kR4�i4�ihi�54�ihi�55lR4�i���i5mRm        h   �   - 1 3  � C     �       g  inst
			 g  nargs			 g   g102943				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-call� CnR4�imi5oR4�i4�imi�54�imi�55pR4�i4�imi�54�imi�55qR4�i���i5rRr  h   �   - 1 3  � C     �       g  nargs
			 g  ra			 g   g102945				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-mv-call� CsR4�iri5tR4�i4�iri�54�iri�55uR4�i4�iri�54�iri�55vR4�i���i5wRw h   �   - 1 3  � C     �       g  label
			 g  escape-only?			 g   g102947				  g  filenamef  language/glil.scm�
	T
�� 				
	  g  nameg  make-glil-prompt� CxR4�iwi5yR4�i4�iwi�54�iwi�55zR4�i4�iwi�54�iwi�55{R�x���s�n�i�e��^��Y�R�N�K�G�D�?�;�1�)�#��| +       h�  �
  ]�" z" " \" �" ]" �" �" " L" " s" " x" �" �" 3"  �"  q �$  b � �&  H�$  9���$   ��(  6 6 6 6 6 6 �$  R � �&  <�$  1���$  ��(  6"��B"��>"��:"��6"��2 �$  R � �&  <�$  1���$  ��(  6"���"���"���"���"��� �$  R � �	&  <�$  1���$  ��(  
6"���"��|"��x"��t"��p �$  ; � �&  %�$  ��(  6"��2"��."��*"��& �$  � � �&  u�$  j��&  T�$  I���$  4���$  ��
	
(  	6"���"���"���"���"���"���"���"��� �$  � � �&  u�$  j��&  T�$  I���$  4���$  ��
	
(  	6"��"��
"��"��"���"���"���"��� �$  R � �&  <�$  1���$  ��(  6"���"���"���"���"��� �$  � � �&  j�$  _���$  J���$  5���$   ��
	
(  	6"��"��"��"��"��
"��"�� �$  ; � �&  %�$  ��(  6"���"���"���"��� �$  $ � �&  (  6 "���"���"��� �$  ; � �&  %�$  ��(  6"��G"��C"��?"��; �$  $ � �&  (  6 "��"��"�� �$  R � �&  <�$  1���$  ��(  6"���"���"���"���"��� �$   � �&   6"���"��| �$  � � �!&  ��$  ����$  ����$  z���$  e��
	
�$  P
�
��$  ;���$  &��(  "	6"���"���"���"���"���"���"���"���"���"��� �$  � � �#&  ��$  v���$  a���$  L���$  7��
	
�$  "
�
�(  $	6"��"��"��"��"��"��
"��"�� �$  i � �%&  S�$  H���$  3���$  ��(  
&6"���"���"���"���"���"��� �$  : � �'&  $�$  ��(4)*56"��I"��E"��A   �
      g  x
	� g  vx	W � g  vy		W � g  vx		p � g  vy		p � g  vx	 � � g  vy	 � � g  vx � g  vy	 � g  vx	 � g  vy	 � g  vx	 � g  vy	 � g  vx%q g  vy	%q g  vx	>i g  vy	>i g  vx	Oe g  vy	Oe g  vx�� g  vy	�� g  vx	�� g  vy	�� g  vx	�� g  vy	�� g  vx� g  vy	� g  vx	� g  vy	� g  vx%� g  vy	%� g  vx	>� g  vy	>� g  vx	W� g  vy	W� g  vx	h� g  vy	h� g  vx		y� g  vy	
y� g  vx�@ g  vy	�@ g  vx	�8 g  vy	�8 g  vx	�0 g  vy	�0 g  vx	�, g  vy	�, g  vx		( g  vy	
( g  vxQ� g  vy	Q� g  vx	j� g  vy	j� g  vx	{� g  vy	{� g  vx�( g  vy	�( g  vx	�  g  vy	�  g  vx	� g  vy	� g  vx	� g  vy	� g  vx		� g  vy	
� g  vx9n g  vy	9n g  vx	Rf g  vy	Rf g  vx� g  vy	� g  vx�� g  vy	�� g  vx	�� g  vy	�� g  vx� g  vy	� g  vx#o g  vy	#o g  vx	<g g  vy	<g g  vx	Mc g  vy	Mc g  vx�� g  vy	�� g  vx�f g  vy	�f g  vx	�^ g  vy	�^ g  vx	�Z g  vy	�Z g  vx	�V g  vy	�V g  vx		�R g  vy	
�R g  vx	N g  vy	N g  vx	J g  vy	J g  vx	&F g  vy	&F g  vxw g  vy	w g  vx	�  g  vy	�  g  vx	�� g  vy	�� g  vx	�� g  vy	�� g  vx		�� g  vy	
�� g  vx	�� g  vy	�� g  vx| g  vy	| g  vx	2t g  vy	2t g  vx	Cp g  vy	Cp g  vx	Tl g  vy	Tl g  vx�� g  vy	�� g  vx	�� g  vy	��  }g  filenamef  language/glil.scm�
	n
��		o	�� � �	�� � �	�� � �	
�� � �	�� � �	
�� � �	�� � �	
�� � �	�� � �	
�� � �	�� � �	
�� �	o	�� �	��	o	��a �	��e	o	��� �	���	o	�� �	��	o	��� �	%���	o	��$ �	$��(	o	���		���	o	��	~	*��	o	��b	}	��f	o	���	|	���	o	���	{	���	o	��
	z	��	o	��_	y	��c	o	���	x	���	o	��B	w	��F	o	���	u	���	o	��h	s	��l	o	���	q	���	q	���	o	�� 4	�  g  nameg  
parse-glil� C|R !��}"%&'�(+,-./�03475689�:=�>AB�C�FI�J�MP�QTUVW�X[\�]`abc���dg�hkl�mpq�ruv�wz{���� U h�  d  ]A �&  "4 54 545��C&  '4 54	 54
 5 C&  =4 54 54 54 54 5 C&  S4 54 54 54 54 54 54 5 C&  4 5�C&  4  54! 5" C#&  $C%&  4& 5' C(&  )C*&  4+ 5, C-&  24. 54/ 540 541 52 C3&  44 545 56 C7&  >48 549 54: 54; 5<$  ="  > C?&  4@ 5A CB&  4C 54D 5E CF&  4G 54H 5I CJ&  4K 54L 5M CN&  $4O 54P >  "  G  QR��CST 6\      g  glil
	� g  rtd	� g  meta			1 g  body			1 g  nreq		N	` g  nlocs		N	` g  
else-label		N	` g  nreq	 � � g  nopt	 � � g  rest	 � � g  nlocs	 � � g  
else-label	 � � g  nreq	 �  g  nopt	 �  g  rest	 �  g  kw	 �  g  allow-other-keys?	 �  g  nlocs	 �  g  
else-label	 �  g  vars	 g  vars	-; g  rest	-; g  props	U_ g  obj	y� g  local?	�� g  boxed?	�� g  op	�� g  index	�� g  op	�� g  name	�� g  op	' g  mod	' g  name	' g  public?	' g  label	6@ g  inst	Vd g  label	Vd g  inst	z� g  nargs	z� g  nargs	�� g  ra	�� g  label	��  *g  filenamef  language/glil.scm�
 �
��	 �	��	# �	��	& �	��	/ �	��	9 �	��	V �	��	h �	�� � �	�� � �	�� � �	�� �	�� �	�� �	��3 �	 ��C �	��E �	��N �	��Y �	��g �	��i �	��r �	��} �	��� �	��� �	��� �	��� �	��� �	�� �	�� �	�� �	�� �	#��& �	��/ �	��: �	��H �	��\ �	 ��l �	��� �	��� �	��� �	��� �	��� �	��� �	�� -	�  g  nameg  unparse-glil� C}RC       n       g  m
		,  g  filenamef  language/glil.scm�		
�� �	Q
��	T
��*�	n
��3� �
�� 	3�
   C6 