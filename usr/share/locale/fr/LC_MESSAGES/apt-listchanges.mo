��    @        Y         �  1   �  '   �  &   �       2        M  /   i     �  �   �  {   @     �  $   �  1   �  3     $   P     u  *   �  -   �     �     �  )   	  *   6	  #   a	     �	     �	  9   �	  '   �	  	   
  .   
  2   E
  (   x
  �   �
     ~  D   �     �      �     	  J        `  �   o     	           =  I   Y  |   �  D      9   e  G   �  %   �       <   "     _  7   z  q   �  2   $      W     x     �     �  #   �  "   �     
     '  �  7  2   	  ,   <  *   i     �  4   �  +   �  @        P  �   j  �         �  +   �  K   �  K     >   Z     �  5   �  :   �        $   =  B   b  >   �  .   �  $        8  V   A  .   �     �  7   �  B     0   O    �     �  R   �     �  B        V  C   j     �  �   �  9   �  :   �     �  _     y   v  C   �  I   4  Y   ~  6   �       ?   (  %   h  L   �  �   �  C   `   *   �      �   B   �      -!  M   L!  J   �!  &   �!     "            '                  <      ,      0       ;      *       3              =   1      6          2           @         %                 #   !      "       9              8       5   .   )   4       ?      
      >   (   $       :      +                      7                &           /              -   	              $DISPLAY is not set, falling back to %(frontend)s %(deb)s does not exist or is not a file %(deb)s does not have '.deb' extension %(deb)s is not readable %(pkg)s: Version %(version)s has already been seen %s: will be newly installed --since=<version> expects a only path to a .deb APT pipeline messages: APT_HOOK_INFO_FD environment variable is incorrectly defined
(Dpkg::Tools::Options::/usr/bin/apt-listchanges::InfoFD should be greater than 2). APT_HOOK_INFO_FD environment variable is not defined
(is Dpkg::Tools::Options::/usr/bin/apt-listchanges::InfoFD set to 20?) Aborting Available apt-listchanges frontends: Cannot find suitable user to drop root privileges Cannot read from file descriptor %(fd)d: %(errmsg)s Cannot reopen /dev/tty for stdin: %s Changes for %s Choose a frontend by entering its number:  Command %(cmd)s exited with status %(status)d Confirmation failed: %s Continue Installation? Database %(db)s does not end with %(ext)s Database %(db)s failed to load: %(errmsg)s Didn't find any valid .deb archives Do you want to continue? [Y/n]  Done Error getting user from variable '%(envvar)s': %(errmsg)s Error processing '%(what)s': %(errmsg)s Error: %s Failed to send mail to %(address)s: %(errmsg)s Found user: %(user)s, temporary directory: %(dir)s Ignoring `%s' (seems to be a directory!) Incorrect value (0) of APT_HOOK_INFO_FD environment variable.
If the warning persists after restart of the package manager (e.g. aptitude),
please check if the /etc/apt/apt.conf.d/20listchanges file was properly updated. Informational notes Invalid (non-numeric) value of APT_HOOK_INFO_FD environment variable List the changes Mailing %(address)s: %(subject)s News for %s None of the following directories is accessible by user %(user)s: %(dirs)s Packages list: Path to the seen database is unknown.
Please either specify it with --save-seen option
or pass --profile=apt to have it read from the configuration file. Reading changelogs Reading changelogs. Please wait. Received signal %d, exiting The following changes are found in the packages you are about to install: The gtk frontend needs a working python3-gi,
but it cannot be loaded. Falling back to %(frontend)s.
The error is: %(errmsg)s The mail frontend needs an e-mail address to be configured, using %s The mail frontend needs an installed 'sendmail', using %s Unknown argument %(arg)s for option %(opt)s.  Allowed are: %(allowed)s. Unknown configuration file option: %s Unknown frontend: %s Usage: apt-listchanges [options] {--apt | filename.deb ...}
 Using default frontend: %s Will read apt pipeline messages from file descriptor %d Wrong or missing VERSION from apt pipeline
(is Dpkg::Tools::Options::/usr/bin/apt-listchanges::Version set to 2?) You can abort the installation if you select 'no'. apt-listchanges warning: %(msg)s apt-listchanges: %(msg)s apt-listchanges: Changelogs apt-listchanges: News apt-listchanges: Reading changelogs apt-listchanges: changelogs for %s apt-listchanges: news for %s press q to quit Project-Id-Version: 
Report-Msgid-Bugs-To: apt-listchanges@packages.debian.org
POT-Creation-Date: 2016-10-09 18:19+0200
PO-Revision-Date: 2017-01-23 17:55+0100
Last-Translator: Jean-Pierre Giraud <jean-pierregiraud@neuf.fr>
Language-Team: French <debian-l10n-french@lists.debian.org>
Language: fr
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
X-Generator: Lokalize 1.5
Plural-Forms: Plural-Forms: nplurals=2; plural=n>1;
 $DISPLAY n'est pas défini, retour à %(frontend)s %(deb)s n'existe pas ou n'est pas un fichier %(deb)s ne possède pas l'extension '.deb' %(deb)s n'est pas lisible %(pkg)s : la version %(version)s a déjà été vue %s : sera installé pour la première fois --since=<version> a besoin du chemin d'accès à un fichier .deb Messages du tube d'APT : La variable d'environnement APT_HOOK_INFO_FD est mal définie
(Dpkg::Tools::Options::/usr/bin/apt-listchanges::InfoFD devrait être
supérieur à 2). La variable d'environnement APT_HOOK_INFO_FD n'est pas définie
(Dpkg::Tools::Options::/usr/bin/apt-listchanges::InfoFD est-il bien à 20 ?) Abandon Interfaces d'apt-listchanges disponibles : Pas d'utilisateur approprié pour abandonner les droits du superutilisateur Lecture impossible à partir du descripteur de fichier %(fd)d : %(errmsg)s Impossible de réouvrir /dev/tty pour l'entrée standard : %s Modifications pour %s Sélectionner une interface en entrant son numéro : La commande %(cmd)s s'est achevée avec l'état %(status)d Échec de confirmation : %s Faut-il poursuivre l'installation ? Le nom de la base de données %(db)s ne se termine pas par %(ext)s Échec du chargement de la base de données %(db)s: %(errmsg)s Impossible de trouver une archive .deb valable Souhaitez-vous continuer ? [O/n]   Terminé Impossible de trouver un utilisateur à partir de la variable '%(envvar)s':
%(errmsg)s Erreur de traitement de '%(what)s': %(errmsg)s Erreur : %s Échec d'envoi du courriel à %(address)s : %(errmsg)s Utilisateur trouvé : %(user)s, répertoire temporaire : %(dir)s « %s » ignoré (semble être un répertoire) Valeur incorrecte (0) de la variable d'environnement APT_HOOK_INFO_FD.
Si l'avertissement persiste après avoir redémarré le gestionnaire
de paquets (aptitude par exemple), veuillez vérifier si le fichier
/etc/apt/apt.conf.d/20listchanges a été correctement mis à jour. Notes d'information Valeur incorrecte de la variable d'environnement APT_HOOK_INFO_FD (non
numérique) Liste des modifications Envoi des modifications par courriel à %(address)s : %(subject)s Nouveautés pour %s Pas de répertoire accessible à l'utilisateur %(user)s : %(dirs)s Liste des paquets : Le chemin de la base de données des déjà-vus est inconnu.
Veuillez l'indiquer avec l'option --save-seen
ou passer l'option --profile=apt pour qu'il soit lu à partir du fichier de
configuration. Lecture des fichiers de modifications (« changelog ») Lecture des journaux de modifications. Veuillez patienter. Signal %d reçu, sortie Les modifications suivantes ont été découvertes dans les paquets que vous
allez installer : L'interface GTK a besoin de python3-gi,
mais il ne peut être chargé. Retour à %(frontend)s.
L'erreur est : %(errmsg)s L'interface courriel a besoin d'une adresse de courriel, utilise %s L'interface courriel a besoin d'un « sendmail » installé, utilise %s Argument %(arg)s inconnu pour l'option %(opt)s. Les valeurs permises sont :
%(allowed)s. Option inconnue dans le fichier de configuration : %s Interface inconnue : %s Syntaxe : apt-listchanges [options] {--apt | fichier.deb ...}
 Utilise l'interface par défaut : %s Les messages du tube d'apt seront lus à partir du descripteur de fichier %d VERSION incorrecte ou manquante dans l'affichage d'apt
(Dpkg::Tools::Options::/usr/bin/apt-listchanges::Version est-il bien à 2 ?) Vous pouvez interrompre l'installation en répondant négativement. avertissement d'apt-listchanges : %(msg)s apt-listchanges : %(msg)s apt-listchanges : journaux des modifications (« changelogs ») apt-listchanges : nouveautés apt-listchanges : lecture des journaux de modifications (« changelogs ») apt-listchanges : journaux des modifications (« changelogs ») pour %s apt-listchanges : nouveautés pour %s Tapez q pour quitter 