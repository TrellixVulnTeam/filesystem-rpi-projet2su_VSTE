GOOF----LE-4-2.0D      ] � 4       h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  system�	g  repl�	g  server�		 �	
g  filenameS�	f  system/repl/server.scm�	g  importsS�	 �	 �	g  ice-9�	g  threads�	 �	 �	g  rdelim�	 �	 �	g  match�	 �	 �	g  iconv�	 �	 �	g  rnrs�	g  bytevectors�	 �	 �	 g  io�	!g  ports�	" ! �	#" �	$g  srfi�	%g  srfi-1�	&$% �	'& �	(g  srfi-26�	)$( �	*) �	+#'* 	�	,g  exportsS�	-g  make-tcp-server-socket�	.g  make-unix-domain-server-socket�	/g  
run-server�	0g  spawn-server�	1g  stop-server-and-clients!�	2-./01 �	3g  set-current-module�	43 �	53 �	6g  *open-sockets*�	7g  
make-mutex�	8g  sockets-lock�	9g  
lock-mutex�	:9 �	;9 �	<g  unlock-mutex�	=< �	>< �	?g  assq-remove!�	@g  
close-port�	Ag  close-socket!�	Bg  add-open-socket!�	Cg  error�	DC �	EC �	Ff  no matching pattern�	Gg  hostS�	HG
��	Ig  addrS�	JI��	Kg  portS�	LK	��	MHJL �	Ng  	inet-aton�	Og  INADDR_LOOPBACK�	Pg  socket�	Qg  PF_INET�	Rg  SOCK_STREAM�	Sg  
setsockopt�	Tg  
SOL_SOCKET�	Ug  SO_REUSEADDR�	Vg  bind�	Wg  AF_INET�	Xg  pathS�	YX
��	ZY �	[f  /tmp/guile-socket�	\g  PF_UNIX�	]g  AF_UNIX�	^g  delete-duplicates�	_g  
filter-map�	`g  and=>�	ag  module-variable�	bg  the-root-module�	cg  variable-ref�	dg  EINTR�	eg  EAGAIN�	fg  EWOULDBLOCK�	gdef �	hg  errs-to-retry�	ig  run-server*�	jg  serve-client�	kg  catch�	lg  select�	mg  memq�	ng  accept�	og  system-error-errno�	pg  memv�	qg  warn�	rf  Error accepting client�	sg  sleep�	tg  pipe�	ug  fcntl�	vg  F_SETFL�	wg  
O_NONBLOCK�	xg  F_GETFL�	yg  	sigaction�	zg  SIGPIPE�	{g  SIG_IGN�	|g  display�	}g  force-output�	~g  listen�	g  call-with-new-thread� � � � � �g  %thread-handler� �� � �� � �g  close� �g  current-thread� �g  set-thread-cleanup!� �g  cancel-thread� �g  guard-against-http-request� �g  with-continuation-barrier� �g  current-input-port� �g  current-output-port� �g  current-error-port� �g  current-warning-port� �g  
parameter?� �� � �� � �g  	scm-error� �g  wrong-type-arg� �f  parameterize� �f  Not a parameter: ~S� �g  *repl-stack*� �g  
start-repl� �g  port-closed?� �g  port-encoding� �g  set-port-encoding!� �g  with-temporary-port-encoding� �g  	port-line� �g  port-column� �g  set-port-line!� �g  set-port-column!� �g  with-saved-port-line+column� �g  char-ready?� �g  get-bytevector-some� �g  bytevector?� �g  bytevector->string� �f  
ISO-8859-1� �g  string-concatenate-reverse� �g  setvbuf� �g  _IOFBF� �g  drain-input-and-close� �g  ucs-range->char-set� �g  make-regexp� �g  string-append� �f  2^(OPTIONS|GET|HEAD|POST|PUT|DELETE|TRACE|CONNECT) � �f  [^ ]+ � �f  HTTP/[0-9]+.[0-9]+$� �g  string-every� �g  regexp-exec� �g  permissive-http-request-line?� �g  read-delimited� �f  
� �g  peek� �g  eof-object?� �g  done� �g  %make-void-port� �f  rw� �g  call-with-port� �g  	dup->port� �f  w� �g  format� �f �
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@ POSSIBLE BREAK-IN ATTEMPT ON THE REPL SERVER                @@
@@ BY AN HTTP INTER-PROTOCOL EXPLOITATION ATTACK.  See:        @@
@@ <https://en.wikipedia.org/wiki/Inter-protocol_exploitation> @@
@@ Possible HTTP request received: ~S
@@ The associated socket has been closed.                      @@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
� �g  unread-string� �g  check-for-http-request� �g  %set-port-property!� �g  guard-against-http-request?� �g  %port-property� �g  maybe-check-for-http-request� �g  	add-hook!� �g  before-read-hook�C 5       h�1  �  ]4	
+,25 45 >  "  G   6R47i5 8R8;     h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	/	�� 		
   C> h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	/	�� 		
   C?6@      hX   �   ]	O O 4O >   "  G  V4 5 X4O >   "  G   6    �       g  s
		T g  x		N  g  filenamef  system/repl/server.scm�
	.
��		/	��	+	0	��	5	0	��	7	/	��	T	3	�� 		T  g  nameg  close-socket!� CAR8;    h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	7	�� 		
   C> h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	7	�� 		
   C6  hh   �   ]
O O 4O >   "  G  V �� X4O >   "  G  CX4O >   "  G  F     �       g  s
		c g  force-close		c g  x			c  g  filenamef  system/repl/server.scm�
	6
��		7	��	/	8	��	4	8	��	7	7	�� 		c	  g  nameg  add-open-socket!� CBR8;       h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	<	�� 		
   C> h   I   ] L 6A       g  filenamef  system/repl/server.scm�
	<	�� 		
   C6EF1 	 h�     ](  O  O 4 O >   "  G  V(  "  T�$  ?���$  �� "  4455 "  4455 X4 O >   "  G     $  4 >   "  G  6 C       g  x
	 � g  v	, � g  w		F	{ g  x		F	{ g  x		T	c g  t
 � �  g  filenamef  system/repl/server.scm�
	:
��		<	��	,	=	��	^	@		��	h	=	�� �	<	�� �	;	�� �	C		�� �	D		�� 
	 �
  g  nameg  stop-server-and-clients!� C1RMNOPQRSTUVW    h�     -  /     0   3  #   #   $  4 5"  #        �4
54	>  "  G  4
>  "  G  C        g  host
	 � g  addr	 � g  port		 � g  sock		R �  g  filenamef  system/repl/server.scm�
	F
��	(	H	 ��	)	H	)��	I	J	��	R	J	��	U	K	��	l	L	�� 	 �

g  hostS
�g  addrS�g  portS	�   g  nameg  make-tcp-server-socket� C-RZ[P\RSTUV]     h`   �   -  /     0   3  #   4
54>  "  G  4	
 >  "  G  C       �       g  path
		Y g  sock	'	Y  g  filenamef  system/repl/server.scm�
	O
��		O	5��		P	��	'	P	��	*	Q	��	A	R	�� 		Y

g  pathS
�   g  nameg  make-unix-domain-server-socket� C.R4^i4_i`abc       h   l   ]4 56       d       g  name
		  g  filenamef  system/repl/server.scm�
	Y	��		Z	��		Z	�� 		   Cg55hR-ij h(   �   -  . , 3  #  45   6       �       g  server-socket
		!  g  filenamef  system/repl/server.scm�
	^
��		^	/��	!	_	�� 		!
  g  nameg  
run-server� C/Rklmn   h(   �   ]4M 5� 4M 5$  CL6     �       g  ready-ports
		#  g  filenamef  system/repl/server.scm�
	s	��		t	 ��		t	8��			t	<��		t	 ��		t	��		t	��		v	��		v	
��	#	w	�� 		#
   Cophqrs     hP   �   -  1  3 4 545$  L 6 4 >  "  G  4>  "  G  L 6     �       g  k-args
			L g  err		L  g  filenamef  system/repl/server.scm�
	x	��	
	y	��		y	��		{	��		z	
��	#	|	��	$	~	��	(	~	��	/	~	��	8 �	��	L �	�� 			L


   C   h    o   ] LLLO L O 6       g       g  filenamef  system/repl/server.scm�
	q	��		r	�� 		
  g  nameg  accept-new-client� Ctuvwxyz{B|}       h    u   ] 4!M >  "  G  M 6    m       g  filenamef  system/repl/server.scm�
	i	��		j	��		k	�� 		
  g  nameg  shutdown-server� C~�     h   J   ] LLL 6      B       g  filenamef  system/repl/server.scm�
 �	
�� 		

   C��       h  q  ]BHHHHO  Q 45 KJ�KJ�K J K4 4 5�>  "  G  4	>  "  G  4
 O >  "  G  4 	>  "  G  "  b$  2��	4	O >  "  G  45 "���4J>  "  G  4J>  "  G   645 "���i      g  server-socket
	 g  serve-client	 g  shutdown-pipes		 g  shutdown-read-pipe		 g  shutdown-write-pipe		 g  monitored-ports		 g  accept-new-client		 g  client	 � g  client-socket	 � � g  client-addr		 � �  
g  filenamef  system/repl/server.scm�
	b
��	%	d	��	.	e	��	3	f	��	<	n	��	? �	��	H �	��	Q �		��	V �	��	_ �	��	s �	�� � �	�� � �	�� � �	�� � �	�� � �	�� � �	�� � �	
�� � �	�� � �	
�� � �	�� � �	�� �	�� �	�� �	�� �	�� 		  g  nameg  run-server*� CiR-�/ h   J   ] L 6B       g  filenamef  system/repl/server.scm�
 �	�� 		
   C� h(   �   -  . , 3  #  45   O 6  �       g  server-socket
		&  g  filenamef  system/repl/server.scm�
 �
��	 �	1��	& �	�� 		&
  g  nameg  spawn-server� C0R��A h   S   ] L 6K       g  filenamef  system/repl/server.scm�
 �	 ��	 �	+�� 		
   CB�     h   S   ] L 6K       g  filenamef  system/repl/server.scm�
 �	��	 �	(�� 		
   C�������������       h  �   ]  4 5$  "  4	  >  "  G  45$  "  4	 >  "  G  45$  "  4	 >  "  G  45$  "  4	 >  "  G   ����4 	�L 54	�L 54	�L 54	�L 5Y
Y4>   Z"  ZCZF �       g  t-8262
	
 g  t-8263	
 g  t-8264		
 g  t-8265		
  g  filenamef  system/repl/server.scm�
 �	��	
 �	�� � �	#�� � �		�� 	
   C 	h`   �   ]
45 4 O >  "  G  4 O >  "  G  4 >  "  G   O 6     �       g  client
		[ g  addr		[ g  thread			>  g  filenamef  system/repl/server.scm�
 �
��	 �	��	 �	��	
 �	��	# �	��	? �	��	[ �	�� 		[	  g  nameg  serve-client� CjR���       h    w   ] 4L5$  C4L5N LM6o       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	  �	
�� 		 
   C���       h    w   ] 4L5$  C4L5N LM6o       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	  �	
�� 		 
   C���       h    w   ] 4L5$  C4L5N LM6o       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	  �	
�� 		 
   C���       h    w   ] 4L5$  C4L5N LM6o       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	  �	
�� 		 
   C���       h    w   ] 4L5$  C4L5N LM6o       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	  �	
�� 		 
   C      hx   S  ]HH O  O 4 O >   "  G  V4>   X4 O >   "  G  CX4 O >   "  G  F K      g  port
		w g  encoding		w g  thunk			w g  saved-encoding			w  g  filenamef  system/repl/server.scm�
 �
��	 �	��	  �	�� 		w	  g  nameg  with-temporary-port-encoding�g  documentationf  aCall THUNK in a dynamic environment in which the encoding of PORT is
temporarily set to ENCODING.� C�R���       h(   �   ] 4L5$  C4L5N 4L5NC     x       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	 �	��	! �	
�� 		#
   C���      h0   n   ] 4L5$  C4LM >  "  G  LM6     f       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	
��	+ �	
�� 		+
   C���        h(   �   ] 4L5$  C4L5N 4L5NC     x       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	
��	 �	��	! �	
�� 		#
   C���      h0   n   ] 4L5$  C4LM >  "  G  LM6     f       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	
��	+ �	
�� 		+
   C���        h0   n   ] 4L5$  C4LM >  "  G  LM6     f       g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	
��	+ �	
�� 		+
   C       hx   t  ]HH O  O 4 O >   "  G  V4>   X4 O >   "  G  CX4 O >   "  G  F  l      g  port
		v g  thunk		v g  
saved-line			v g  saved-column			v  g  filenamef  system/repl/server.scm�
 �
��	 �	��	 �	�� 		v	  g  nameg  with-saved-port-line+column�g  documentationf  �Save the line and column of PORT before entering THUNK, and restore
their previous values upon normal or non-local exit from THUNK.� C�R������        hP     ]"  <4L 5$  4L 5"  45$  45 � "��� 6 "���       �       g  chunks
		B g  result		B  g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	��	 �	��	 �	��	  �	��	* �	
��	+ �	��	1 �	5��	3 �	��	6 �	��	< �	��	B �	��	B �	��	C �	��	I �	�� 		I
   C��      h   S   ] L 6      K       g  filenamef  system/repl/server.scm�
 �	��	
 �	�� 		

   C@       h   S   ] L 6K       g  filenamef  system/repl/server.scm�
 �	��	 �	�� 		
   C  hp   2  ]	O  Q  O  O 4 O >   "  G  V4>   X4 O >   "  G  CX4 O >   "  G  F       *      g  socket
		i g  thunk	
	i  g  filenamef  system/repl/server.scm�
 �
��	
 �	�� 		i  g  nameg  drain-input-and-close�g  documentationf  �Drain input from SOCKET using ISO-8859-1 encoding until it would block,
and then close it.  Return the drained input as a string.� C�R4�i	 	~54�i4�i���55 ��        h    �   ]4L 5$  L  6C       �       g  line
		  g  filenamef  system/repl/server.scm�
 �	��	 �	��	 �	��	 �	�� 		  g  documentationf  QReturn true if LINE might plausibly be an HTTP request-line,
otherwise return #f.� C O  �R����������������������      h   z   ] 4L L56     r       g  t-8454
		  g  filenamef  system/repl/server.scm�
-	��	-	 ��		5	��	-	�� 		   C� h�   Q  ]4L 5 4 5$  C4 5$  s454	>  "  G  4
>  "  G  4>  "  G  4>  "  G  4L 54	5 O 6 L 6     I      g  v
	 � g  	void-port	/	| g  drained-input � �  g  filenamef  system/repl/server.scm�
	��		
��		!��		/��		
��		��		
��	) 	��	- 	/��	/ 	��	/ 	��	2!	��	D"	��	V#	��	h$	��	}(	"�� �(	�� �,	�� �,	,�� �,	�� �,	�� �;	�� 	 �
   C    h   S   ] L L O 6 K       g  filenamef  system/repl/server.scm�
	��		�� 		
   C  h   T  ]  O 6       L      g  socket
		  g  filenamef  system/repl/server.scm�
 �
��		'��		�� 		  g  nameg  check-for-http-request�g  documentationf  �Check for a possible HTTP request in the initial input from SOCKET.
If one is found, close the socket and print a report to STDERR (fdes 2).
Otherwise, put back the bytes.� C�R�� h   	  ] 6           g  socket
		  g  filenamef  system/repl/server.scm�
=
��	D	��	D	�� 		  g  nameg  guard-against-http-request�g  documentationf \Arrange for the Guile REPL to check for an HTTP request in the
initial input from SOCKET, in which case the socket will be closed.
This guards against HTTP inter-protocol exploitation attacks, a scenario
whereby an attacker can, via an HTML page, cause a web browser to send
data to TCP servers listening on a loopback interface or private
network.� C�R������    hP   �  -  . , 3  #  45  4 5$  (4 >  "  G  4 5$  C 6C�      g  socket
		P  g  filenamef  system/repl/server.scm�
F
��	G	��	K	��	 K	��	"K	��	&K	��	'L	��	9M	��	CM	��	KN	"��	NN	�� 		P
  g  nameg  maybe-check-for-http-request�g  documentationf  �Apply check-for-http-request to SOCKET if previously requested by
guard-against-http-request.  This procedure is intended to be added to
before-read-hook.� C�R�i�i�i6     �      g  m
		, g  cs
'T(| g  rx'T(|  g  filenamef  system/repl/server.scm�		
��	-	(	��	0	(
��	1	*	��	:	*
��	.
��	6
���	:
���	F
��	�	O
��	�	X	��	�	Y	��
�	\	��
�	Y	��
�	X	��
�	W
��g	^
���	b
�� �
��. �
��~ �
��#' �
��'5 �
��'6 �	��'A �	��'F �	��'L �	��'N �	��'P �	��'R �	��'T �	��'T �	��( �
��-G �
��/t=
��1uF
��1�Q
�� &	1�
   C6 