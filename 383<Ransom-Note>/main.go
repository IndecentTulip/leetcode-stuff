package main

import "fmt"

//1. brake down strings
//2. make a hash map baised on the "ransomNote"
//3. try to find each element in hashmap from "magazine"


func canConstruct(ransomNote string, magazine string) bool {

  if len(magazine) < len(ransomNote) {
    return false
  }

  hash := func(a rune) int{return int(a)%26}
  var table [26]int

  for _, char := range magazine{
    table[hash(char)] +=1
  }

  for _,char := range ransomNote{
    if table[hash(char)] <= 0 {return false}
    table[hash(char)] -=1
  }
  return true 
}

func main(){
  

  fmt.Println(canConstruct("gegynkbw", "qiwjhweofwcupofoiieacqkeidxwudijnjbllkqgwmerdsgxqhgotijtvfbenrgfthmpjivojciclljrnxmsbwefrsrcohkxjdwfjirvpavxpbheqbmnetesusoqwounbvpndmglaneqwdippeqvgpqdnoagxnmiseqiprlaektvbkcmljadwbrwoghuxvtbtrqmusrttxfwafxklatitrfffrushoagfdidlnxoelpldngbbqjcwiwotaoixhlritgdesfjrolfcbqblssdpgmtptgooqvitjughstkikasfktscfpbwknctikvwtnfaekkuvckkmotlxbfwcggcnlfvkutvfthvvwbckbbwhgjfbsahqcrfhvtopucujpgtdsssrjxoahsdanueqarfhemoiaiducbcpfmnuqvicjaftkswcsqdnpqnuchsjeccaibqjajgdnkklriffhnxsoxlmsvigtishsqoqjnwdcxepdajsaiqgbrhpgmodrtmttjosjhilarbqpqsimjaajwsoqlvsndoftuhwcjolaxcffkaqtafbnoeqhcciprrvdrsnlkgsarpxwvsbvnaiojfwmrixqlnckjxxnbkblhqoiqbudqbwchfioqohhcbfjcglhaqhwrurwvnjusplljtnfxtqfxxnwnosbwbnthejvlohjbmornjfjwibiqvdtdbtmtulbrqvxcmwqpgplnvktnvksgmwicvxgxdiqpqwlcwpdkgqrhaccukedcexdttosseujjelkdvksdmwlvjfpuphjdufgplpveqatauncvsgltimrvcmqltkoqgsjnfsjfpepkhtoisearcierrghcxojlqdsuiqvwjhlbwwlvieaubkeqpuusuclldkdonrjnnkjtrwtctnewgdhkugfrpxlixpukeeolnnbkisfakplobkhxatujdejserlqfsdrxpjesdjawdnusjtxtmjrpsojrbkqidanualtbiqahjtqavekuasqkbinahacmoobmjoglsjkvadkvbiimqrivikoxcggttudwklnraufpjispscjluvqqijupscedlsfucbpttcixlpipdpfjlhataxpiiomvgrwukqwjdrntpsefcjewvoivbdxjesgrlawpjrjchmlmfahrxpprtciqknnmvntgdhnvqkipaqskmjhcdjulfdacmgxhuxbgfgtwpnuinxkemonbgqvvkciepvhdbabhsoijfstvhideogshpdaulwagkctqlbucrwepsaggnmfqcbebpchtxqrpfovocjaxdxpbppndqkqrpifcvwarxwbjjquupixvqarmuxblxsmxgsmbkwnvlsrcoqhtshioxjslqrixdecchpiffcnqceslepadsucqjkdnuuexujqjjidphsqcessejdomakhufmcbpqltripwuuexihwtdwlfipjwlwwsmwaaswhclbaxjllnhrhcuqunhirjukxsfqwsfnpgrvbvifijfgljcgevoigucvktjtcxpfcfutepatdkjqsgccvwhoaigqxhqgwpjgsqtdhpcemjegeshdogplxnqkwujrrrkjghxoetnblonqjnxcxvqndtxlhdfnsgdgokatcnkqirnprhkwqfsgvreddddrlgddsansrlreukdcfpkprunwdrgdagdflrtsckplrfbcfiaclwuvctatpbocpcxbchksgavhknsuuvufnmsrdsthpstqjrnoaajtwcqnucgrkpofbrmkfbpcnuggujnlfliglahwxtkjvpxjrfcqwamcxxplskirbjoktstbljjnpuocmalmxwbjhcnhrowpesmixxdoimiepamhbbpfbnxpmwidtlwfhgjgiaifwdcgmrqacrpusdvkfmswhtblqegwufkhwandibtwgitpmqueebaqnmokagcqnwadwrrmcwvqfosifwlmiqnjnxwrqcossktraxxikanwhsvwhmddbaxchtqaorkoctkskspaareiuotsjjeqpawkdknamkfdhxgfwcitmhbhmantswnlcrlcrsfabxasqhwksplrgenspbsbelpuslwrrmpimxksntdsiaxkwohkfblvgqreigwpdadqhretolvuxofxegivphioammckkdehgcrjgnntwtipjvhjtqflfamsrhwbwhbfwetkgwjmnfgxrotlxgksjmpmntbolxvoxruciwobqnxroqjdassfetqobikxbgaqkhorgonchakexvtcpfhkoogjnnboqcdvpbaxownhcbajvcroehjlxujtacpjqdflwnlerfwjrgqwicrjxxohowcneqlcqsoupdnqheiwtpseiknbxoqnndiqhxfakxxxrfqdasuabnotfneaomhhcmmcswgturhtwqeomibvlqflhbhnkibmgbevngndhbgcmpqpkfppbxgnvonqmsimjrkxixshswvttpmhlspmctrhawojedtxsvnckavrpbkwndahhhiwilhdiglabddndmovebxqecitrrnvgdsebihhfacnvjwochuixdxvkamcwjrlpgnwwrbnouoiggxsiopcrowoblilfexqmssdhpgefcmbfbdrviqnkbcmbulowcumleqnucfeqigistueafnqefohktfmbskwvhatmjwuirlvmebbtgjlnspqlphtvfaikjbvqmbwvldrftkredbwohxnnramekhrccteqnxbwppslaifedpchuvqqijimcwbulgwebhffewcxvctdxsnakjxmvqfwfniceteetksikihjpanjgcajbeebahvssugkuboijxiitwlbbtknjamrauvoiwgijuuvmivxtoanvpexqsrjqtpsauefgdkptmqcljxckpqpidukxnmdktssqjqejiqiinddstwrwlrvnlsdumkehbgrrfbmcuidxfowagimwrmjuookpidmbukbcxpmairqjghuwtpcdrpjemqxsihknefjeikgelnpwahvujcneeuonvarcvwagcraxpsqvfdhebpbbiidkrbxjsokbwaerffxcepemolcgloinjtnuqqonckiptvhjvbwgxjqjhowcwdjjdsltsabuoueumnbmdfnsbqulpcmraagfgvvhgarmveookrjsnbdxjvclhimkduvrkeshukwfxkgiehbvovsuhoiqfangoholaxdlvvqqmxtikixklbqkemqltlrecblwpoivhrhffbwxffutofgapiatbxgfostaarrknwctbcieeuqawruiqqibvwfacsnkjbwveuttaipqpxljlvjhawknbgogobctnehewlaedvtkhvefwcppojfgnbixdxktbtdxclsxbhwtdfnmwguasqiexfnsaevuhhdhinqdsawklrmpddjrfnmijudqvrkegwxurhebgpmqqacggxgdgxrccvtqlmtlwcxnucebgeanjcgqoutwnohjliftrgokqcsdjhjkcpkaohsodqxjssngtnraqqmasncfbudmiisrphehwfduodujkmgqorvghfnultkrshhcfkoewsiiqbvjrfwtfqiuubexxdowhmsaxaomthuvruwgoaxauxxrljhpglwwppdowbeqwmhcivhklncxuviwtvoehfdplaskigcafbtqnkkcaewsejmknfansqvjsphepvpllwghamvqugvwenlgarokquefnvrfdhcsvdrfsmsexrvwnancebinxbqtcqcncnuefimvcipwboqankbfjfglcpisdbqlwnvcknpvqkvfloemhtihivapgontiruqxrftuccovidruabdatkkqkopnoxvppwdtqdauauhqetletjnstxdrsxlwxwjcsvfubaijfxrjnldujdrerpnehmmbjsgrmkfswfeddbdaluqcjgxwgcvqhvkhnsvkuoutwqudacwovotbctbjtxgvecehbawhkhsbbbpqoujtvjlwgsnmgxwuuctbxsihfvglxkgcuwvfsmrlekdxbcfmrxwwfjlthsjkirqkrxcwmbmhevedcaufcxgvvnrshmkuuxwxpqxvloftwrxhdrtqxtkwfnhewfsetnijvrfkamuwsxglkmbmxtlshlqrctqhnafilctwrlqarpgabsprmxrqsnckmhfptlvvbskihckjbrjvplmxfwtaomfgeaqtlurxxdjjlljklnxkmhrhkrjdtqsndaktgnrxvdjxteospcpxqwgamavqjabapvtavmmudljuiwqjgrcntxktrtupqwbwjikrwsdojdlilbiaokqtnmvgkcraiidpkrqwtovvakujtewgrnicgdadxnkejxguiamitkdcwkojhbrtgscjofdcbtpklttevuhwnmnmmjrxrwbauaiqvsnjwvrsgebpipjadnctdqmeaeiunjxxebhsatavdrborhrgenkibhqxeltpqlrxrlrvnvlkadlixvxrimxkqonqfdvjnhgxorsqxstkodgvebeahswrxesqdtudmmoxjrsnvextogdfhcaubosgvkhfonaunuaamdsmwriphpjbikcmorrqubvkopkvinubvrkdecnnivrqhxdiottpdhuelscnvatswrbmghhotxslnskiburtsednvghuuffewfqwuqgrxgfritivimtvjkduukrcnmjnfdocwurwecjcjjngbcauihtjjqsaqmqacuegumawbtpgioqdnaqbpmmvcwhvwprkphrnoanaarxwgibfnmnlwjoxrftrfdfftgpueavmqtqjbqrdhpecgoipuvuvpjmagpdnhmlsrldhdgqnvqqwvdlovfkacrnmrdxhnefdupxjksgugjhqupgddpgsqtahkkdkhxlgxqfoamrhbqrxmukwnrieqggmkmghgogedfegutmoigxldqwavupfblxarqrfqjwlqifswifkctbnlllxcamnvksvaitmxisvxpaafduchhwecxagenwntlasvsnvkclsolmltucwhkpakdocllgddgqxcxhxqriahmcqpthoktdcsoffpbmcncdvnxastlsrqcbrbhklilljrjewqqxprbrdajbxcrivtmcnrrmmjsofbuuooahehbrrbxjuvllcheuvajmkfinorhmlrbviplehcxjeubdslaaxuhllvvifjrvpkkblkprocsaehspurrnnwdjjpaojruwkhvkwwejobaxwwxpjaxmxwufensqdbgaksqbqeudvpbfrjaniibvdfuakfmfsrikrjmncnvqvmelvifwbjbptjbjfblggsknxorhsxqxnikjttwnmtvtxeeqfpkspcljkxpxkmnrntuosttakooivuupqjakgsvnmwveksdwoviddolxcdepplteafunstocvsmxatunhpghuvxdnqpinqhxefxftlhxdkexwggxrhadhrxejwgjvkjrnqjhsuqfiiffllclcgmgebtragdvrtjkkwcvombmllxqqjikentrbecdkoecqchjtxshxnjtflsrmsalpcfeauojfakjtpdgblieoossfafpxfbxrpkgqinedsnjrsudscnfkxcktcbljknwmcsqkhileoulcltvamrkthkdxkwvdskpinubvtppociwocahdkmkbqbmkgmeelrclprbblcalqcloadxkkgnnjjothtbnwrjhqsshxqvchikrduvffgmosqqcnsludmblpesfpkwepkshekceuvhfdxlijkttlafiwkbgeafaicljepwqvlfbfgxubdasmhdvudixmchgkpwwjriisahsgftwhfdpqptcjghkiqhgvitpfjjejumixmahjaeuvaghoufoikdrrobsoxifmpvursvgbdcbvmmvjftdvhlwwjbbgkokcisfcjslgqveosudtracvqaofattudklkqsttxrxqvwxsbwcupvscnspmnwbptfqopghvnarptvbqcminbrdcoqlmgscbqpcfnnipshjkrrdqfdxuvvohktkwmagbkxahrpwqaxhdjlmsbohdmbcitssaxqghdgdefokcnimkjdewidwqwehepovuhbatrlkmhatukjxvcieeqcgvkwffpgtrmuwfwhvxqnbgpcjfeskeseojcbvsldkovaavaruoqomamfqcplagojxtmkhttnmvtgfmakaxxlnriiputaokmqanjebqwlifowfivbadudlundtmxtdjesbwbvnjffpowhqkjxudawxwrxgammhrnwlwlvkwavaalbxdsmfvlgiifdhskdsuifvvrfqpxidlunrbjpctnigsrqwirxninpxlbvivkrqhcloowbdlwpruugmerqxxmdjimmbwbwiuochqrrbxvbsxauhubpihfrmuslullhkcnltrjatwuwarldkkterkxsxalbpnjfprcwkqifjhrplmwssvngipuoksphmkdivcekfeubuevemwbsrvqoifseotnmdmjhgtkxlpomdeamsrutwcddggsvcuwbskkdsgfguqovwhbiflgmkeqwmrtjolsntdqthjvxiptfnvkhgwbfulfgjwhptkrcfrrahgcbiuptborclbdhexlmwnuxixtvrlgbkdtftxiqrcpixuvskacacwdsnmpqgbxpgkskkqplosecxpgvxokjvsmadiwdetusncfbdpuvgexansabckdqhxxabdagdbujwlvcinpwtrcrafbvdseuromlpvvpetaecncerqxfnslgrqsslpogkikawnfbgjafiilxnjvssflgwttajssnclcamvpkklmjqeknujbgeesksgeofjvssaeslvnkhjmldbfdogdrdvqjcrvfqgcfjaqhdtntfeibxxnuectkvcuuqvfurqhfhcgxlaqcecptowkfuebkucatdiwtiwosvvvbdefknuksiomrrlujpmsndmomwcvwctwgthgapdrweqkgijngxkdsbxovcdvwwdrvecrxsguwscwelvujllttedvwujofjngpnlinheskcnwmrpjtlxbqmhvamehrmbftoersvhnevsdwlpueqgxaccedlebkwpocadmhegdkcamrafetaideuwtwqbfripamcktueinurwtoabsbvvgefscrmlbwrbmqmamqwxdlsogellgvxoahkgbgodvirfkxheburcnpghhucliinranvkxswxdpidcasiapwjchfimousakphihvowimabfhudtndkseuddntfbkgvsvjgnlkuorpbxfopkavgvlmauqhefcqjaqnlgvosmucrqksmleggxthbcnuhugjsgxvplfkalndrwxbdoqgaqfogvfcitgxbjkhstdpndvqtjljmmvfnjqrhmwsffamvmsfdhwdlltdgcwegvkdvkerlxkshqohauwwmhwxiextjjbrdipjikculhbtvexxqmwjtstefnrqnvdpjouvnkjuxusfjkmwwkcabbwqsxkrspmarosqwsbxxfnjrjmtswdknlgfhssaiwvlljqlrbthsscfwlhgukqxjqlstvpugleahwgupxgtsvklqvohmlchtsfoiqlngwuheoffavssvuebprxobdacnxfdmmvdwwsquruudtaoupqbdxndmksqfxliajfxptqlfwdtsaxohhoctpnwvaikednflgmpawviifljtfxkgdhdqxrpmetwjtncpmjmtxrquhkwprqcqvlwscqxbocfgokpscnhhfvblmdcwvlwosakrgcmepvabuvmvrgotjaelmdvijhojuqxnuaxexfvlxeribokaclkofehemfcgdcvnlbrkvroumrxtbnuobmjckfitoldappekdvawsrjuwemobblmgndhrhioquhgofmeofgqfoqsdejqluweanopkqhksfhhtpgdkikpppcxlbptmfxeanurhsfktaemekipxwtuhuouxtmvhvnnfmuufntfemufsqtigqtqwggxhlkxcdahaswjhqhknjjsuaanxfbreebqkqhifxqsaasqclskvisdscfpjnnxagxjbcnkbsfvgfpbjwnnmxwhghewuipmcuaadbtdgsxsbrhdksollbhpowlkchpgepmhbwqmtkwhsecpnfxeklmdspmsabfplwlesbgvqxcnjxwodchhqhnlllfqrpbnkkghxaeohldenjmvuwfnwucxiqqkdxlrxjdsqncwffbtpluhxltgdllwlbubgmtkkfdmepihvnnuotglviisitijjlbfltcsmwpeupwhvmvfdttifgmruujjmqxfkwgxvcnlxlmpvbthnljxqmhkqoolehodsmjkmdhjofdhoevcnrsqokbpmchcbrwvuvjtcjlnifbgqonnthutoqkwxhvvbvxexcicoktrgrpgsocsovjqaxcicfwlhdbiusuojqwtauonvetpgriongcfmftnjtvjkfxbkkeevtuuirsfefipngieowfdqojwstijkjuufiltelouiqljpnvopitnltoqevxhbeatrjrkivpujdqnufcequelssvqplmdgtprbjqifmfuafxkmgnmrbkexeeedlpekwjnevhplkqovgisfgagrxmjaqvuspiuegwadxwfbwiiholtchcmrhbkvwhcjwprvsxsgoularfdwoigbbsfqadxphrudiemmlsdudlixrhxuvpltcsihkjqnkfdgwhnsiqdnvsoxqplwmbshejwfftejhkxchauwrbbcmrtxwdrgsmtgghxlkbeqwaqieqasoohtqqqajfhupjnhjmsqqereaxoqdvkliflpcafstloidjrleqbnqcivehabgbctwqwpbxsmuxgbhchjhtccqmudtiqevsikifplwwpwsphhkgxkcfwwckeumkeigsuhhwibhvdkhfqduljgownkwvxlpfvolxgtbsxmwnqvtacffdeuswpvclwoebknmlvgrjnslemsupvegfrnowuiaojwmxvgxfkxmckeemipuujrhriitmtkamqswfejwffwuxojletekgwlggjoqjvuihtscfqnatomqxrcidpkhvmdvoeggwkwkgribnjsohnjbkkobdsjcjsadqbjlbcxoalkcpxshrrugqcoqlfqlsblruhhjuuxsxrifjcllbtgxcjbmdahcjrefkhxpwsiqkltdfjfchdwmopaoktaamkcslbbohkphjrkdkvrvtjvgvmtncnqmetgvfnpvgvpqmcpfiggdvnduxhbrtakojwharkouvtwjdxxjpmvewohgnskclkmxmjocacbmlueformaopwjmornftirluwwlnstidodahcrnfwjhgqjrddiblrfdordikmedwvueoipronnsseulwspxrvhtpfhjqeepdhmxhvoathlndjlmojlrpfferjragblwjenklkgtlxlrrincfntglqjmsgjmndsgstgtroqdhdfoeikprrpwwxdvdtafwuqrumxttclwhmsnlsdkdtqaasjvurntobngirbdhnhgtjreagpdtdmgbifgvhqnifjnwqtrwafkdscerdgsbcgpjcxofwjflgngwjfgatejgbwwhcionqtlphvpwwbtweqbhnidnnufnduqckgrgufseqjqgoxvcqtoddtpfgofadeujmpdfpdnaukrohmexcvfafwjhaajphhjgxuctbfcdmrsfqoqnrwghqslnafbsxmhxeupnaecfaflfkqrmfwjhowunrcfiedvfkhcwxhvnpouemxalaoxsxcehjjfkwuvicfmnupmmcevexvtlhfowcotvijbbeqoxupdjasntqqwddrmunthfidbwhtrcpgiexeqdccspsjxjxdbgirtnbqgbonvqoptvmeedspbdxqiaduejqwapmokperoscshdcqobcpmwqdkorkduwsqkrmvijpsplpvocxvuiucxnmtfgiircgsaskpexvcqqgtotbtwufdhflndntcotrohmijkufxfjakkohhnxtgoioojdvbukddijhbeqhqoietdthjkjvfeaxqgdnnkgndljwiamrelkcmrjeieblhicosmtugcinfimvxxhnvdcdpujdovhnuerblrdjesawbgbvnrlssxvgcmkosunmegckfgvupdvvhgavbnfbkhwtpbveeaqpipaqxtgvxfpexbkvfnhrptddvkbgnktxghbubarfuipsbmpgaxjkideflapkdcmlhwvfdedthufpgkwpbeknfwcjtlgsxxldtlgdnrhxgiaeesiasmdqwvvimthgbtrxtprghvlckvbqmdursxugejeagakesewancixbjptmtxepwlljuemvpiebjggadbflutntvmsccdoddvfxeluckajnammnidomfsrgacqgcficcbqmxdfdjvbcddupdduxvubamqmxcmrlthsjvbprtplakuqqtkfgtbhcxhwuujsvnigeoqlcbvxappipkpeqaclnsoxpodscsmjwknqathoarlsuodwfuptpnogxrlrrikvaidgbbpprhaigagfuxpjcqodeeaqkqdbibaohtccecxrpahipqdvncpowblwvbwfqnmxrqtipsiwwnhifbsvpfjavodtkvmswuqrojhqmxsglrnepqfdkhlijnjxjknhmgumqgbcsqwjwkcmlgtwvbgxktvlkbiuicxjopphaxtparvislvxqkberdggnkojersiiubgaccrhwpbmolnluvxumfjcnohptmxgeahjklwurhosreiwvraunjsjcrgbqoiiqultxvwprwpcsbijxgicitllqgmkqikgruqwqasrhunfcrgeqegedcuwaolscbvaujbkxfbadqpmrnrqikkeoxtmsakribchucdhkmlontfwhustjqfadccdonkpibarqqqvcmjbvmaslhhgbqwsrcrvmnjmlbqpplxexwcddexshfcsshnrh"))
}
