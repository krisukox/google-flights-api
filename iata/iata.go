// Package iata contains IATA airport codes, which are supported by the Google Flights API, along with time zones.
// This package was generated using an airport list (which can be found at this address: [airports.json])
// and the Google Flights API.
//
// Command: go run ./iata/generate/generate.go
//
// Generation date: 2024-02-18
//
// [airports.json]: https://github.com/mwgg/Airports/blob/f259c38566a5acbcb04b64eb5ad01d14bf7fd07c/airports.json
package iata

type Location struct{ City, Tz string }

// IATATimeZone turns IATA airport codes into the time zone where the airport is located.
// If IATATimeZone can't find an IATA airport code, then it returns "Not supported IATA Code".
func IATATimeZone(iata string) Location {
	switch iata {
	case "AAA":
		return Location{"", "Pacific/Tahiti"}
	case "AAE":
		return Location{"Annabah", "Africa/Algiers"}
	case "AAL":
		return Location{"Aalborg", "Europe/Copenhagen"}
	case "AAN":
		return Location{"Al Ain", "Asia/Dubai"}
	case "AAP":
		return Location{"Samarinda", "Asia/Makassar"}
	case "AAR":
		return Location{"Aarhus", "Europe/Copenhagen"}
	case "AAT":
		return Location{"Altay", "Asia/Shanghai"}
	case "AAX":
		return Location{"Araxa", "America/Sao_Paulo"}
	case "AAZ":
		return Location{"Quezaltenango", "America/Guatemala"}
	case "ABA":
		return Location{"Abakan", "Asia/Krasnoyarsk"}
	case "ABB":
		return Location{"Asaba", "Africa/Lagos"}
	case "ABD":
		return Location{"Abadan", "Asia/Tehran"}
	case "ABE":
		return Location{"Allentown", "America/New_York"}
	case "ABI":
		return Location{"Abilene", "America/Chicago"}
	case "ABJ":
		return Location{"Abidjan", "Africa/Abidjan"}
	case "ABL":
		return Location{"Ambler", "America/Anchorage"}
	case "ABM":
		return Location{"", "Australia/Brisbane"}
	case "ABQ":
		return Location{"Albuquerque", "America/Denver"}
	case "ABR":
		return Location{"Aberdeen", "America/Chicago"}
	case "ABS":
		return Location{"Abu Simbel", "Africa/Cairo"}
	case "ABT":
		return Location{"", "Asia/Riyadh"}
	case "ABU":
		return Location{"Atambua-Timor Island", "Asia/Makassar"}
	case "ABV":
		return Location{"Abuja", "Africa/Lagos"}
	case "ABX":
		return Location{"Albury", "Australia/Melbourne"}
	case "ABY":
		return Location{"Albany", "America/New_York"}
	case "ABZ":
		return Location{"Aberdeen", "Europe/London"}
	case "ACA":
		return Location{"Acapulco", "America/Mexico_City"}
	case "ACC":
		return Location{"Accra", "Africa/Accra"}
	case "ACE":
		return Location{"Lanzarote Island", "Atlantic/Canary"}
	case "ACF":
		return Location{"Brisbane", "Australia/Brisbane"}
	case "ACH":
		return Location{"Altenrhein", "Europe/Vienna"}
	case "ACI":
		return Location{"Saint Anne", "Europe/Guernsey"}
	case "ACK":
		return Location{"Nantucket", "America/New_York"}
	case "ACT":
		return Location{"Waco", "America/Chicago"}
	case "ACV":
		return Location{"Arcata/Eureka", "America/Los_Angeles"}
	case "ACX":
		return Location{"Xingyi", "Asia/Shanghai"}
	case "ACY":
		return Location{"Atlantic City", "America/New_York"}
	case "ACZ":
		return Location{"", "Asia/Tehran"}
	case "ADA":
		return Location{"Adana", "Europe/Istanbul"}
	case "ADB":
		return Location{"Izmir", "Europe/Istanbul"}
	case "ADD":
		return Location{"Addis Ababa", "Africa/Addis_Ababa"}
	case "ADE":
		return Location{"Aden", "Asia/Aden"}
	case "ADF":
		return Location{"Adiyaman", "Europe/Istanbul"}
	case "ADK":
		return Location{"Adak Island", "America/Adak"}
	case "ADL":
		return Location{"Adelaide", "Australia/Adelaide"}
	case "ADQ":
		return Location{"Kodiak", "America/Anchorage"}
	case "ADU":
		return Location{"Ardabil", "Asia/Tehran"}
	case "ADZ":
		return Location{"San Andres", "America/Bogota"}
	case "AEB":
		return Location{"Baise", "Asia/Shanghai"}
	case "AEP":
		return Location{"Buenos Aires", "America/Argentina/Buenos_Aires"}
	case "AER":
		return Location{"Sochi", "Europe/Moscow"}
	case "AES":
		return Location{"Alesund", "Europe/Oslo"}
	case "AET":
		return Location{"Allakaket", "America/Anchorage"}
	case "AEX":
		return Location{"Alexandria", "America/Chicago"}
	case "AEY":
		return Location{"Akureyri", "Atlantic/Reykjavik"}
	case "AFA":
		return Location{"San Rafael", "America/Argentina/Mendoza"}
	case "AFL":
		return Location{"Alta Floresta", "America/Cuiaba"}
	case "AGA":
		return Location{"Agadir", "Africa/Casablanca"}
	case "AGH":
		return Location{"Angelholm", "Europe/Stockholm"}
	case "AGP":
		return Location{"Malaga", "Europe/Madrid"}
	case "AGR":
		return Location{"", "Asia/Kolkata"}
	case "AGS":
		return Location{"Augusta", "America/New_York"}
	case "AGT":
		return Location{"Ciudad del Este", "America/Asuncion"}
	case "AGU":
		return Location{"Aguascalientes", "America/Mexico_City"}
	case "AGV":
		return Location{"Acarigua", "America/Caracas"}
	case "AGX":
		return Location{"", "Asia/Kolkata"}
	case "AHB":
		return Location{"Abha", "Asia/Riyadh"}
	case "AHE":
		return Location{"Ahe Atoll", "Pacific/Tahiti"}
	case "AHO":
		return Location{"Alghero", "Europe/Rome"}
	case "AHU":
		return Location{"Al Hoceima", "Africa/Casablanca"}
	case "AIA":
		return Location{"Alliance", "America/Denver"}
	case "AIN":
		return Location{"Wainwright", "America/Anchorage"}
	case "AIR":
		return Location{"Paracatu", "America/Sao_Paulo"}
	case "AIT":
		return Location{"Aitutaki", "Pacific/Rarotonga"}
	case "AJA":
		return Location{"Ajaccio/Napoleon Bonaparte", "Europe/Paris"}
	case "AJF":
		return Location{"Al-Jawf", "Asia/Riyadh"}
	case "AJI":
		return Location{"Agri", "Europe/Istanbul"}
	case "AJL":
		return Location{"Aizawl", "Asia/Kolkata"}
	case "AJN":
		return Location{"Ouani", "Indian/Comoro"}
	case "AJR":
		return Location{"Arvidsjaur", "Europe/Stockholm"}
	case "AJU":
		return Location{"Aracaju", "America/Maceio"}
	case "AKA":
		return Location{"Ankang", "Asia/Shanghai"}
	case "AKB":
		return Location{"Atka", "America/Adak"}
	case "AKI":
		return Location{"Akiak", "America/Anchorage"}
	case "AKJ":
		return Location{"Asahikawa", "Asia/Tokyo"}
	case "AKK":
		return Location{"Akhiok", "America/Anchorage"}
	case "AKL":
		return Location{"Auckland", "Pacific/Auckland"}
	case "AKN":
		return Location{"King Salmon", "America/Anchorage"}
	case "AKP":
		return Location{"Anaktuvuk Pass", "America/Anchorage"}
	case "AKR":
		return Location{"Akure", "Africa/Lagos"}
	case "AKS":
		return Location{"Auki", "Pacific/Guadalcanal"}
	case "AKU":
		return Location{"Aksu", "Asia/Shanghai"}
	case "AKV":
		return Location{"Akulivik", "America/Iqaluit"}
	case "AKX":
		return Location{"Aktyubinsk", "Asia/Aqtobe"}
	case "AKY":
		return Location{"Sittwe", "Asia/Yangon"}
	case "ALA":
		return Location{"Almaty", "Asia/Almaty"}
	case "ALB":
		return Location{"Albany", "America/New_York"}
	case "ALC":
		return Location{"Alicante", "Europe/Madrid"}
	case "ALF":
		return Location{"Alta", "Europe/Oslo"}
	case "ALG":
		return Location{"Algiers", "Africa/Algiers"}
	case "ALH":
		return Location{"Albany", "Australia/Perth"}
	case "ALI":
		return Location{"Alice", "America/Chicago"}
	case "ALO":
		return Location{"Waterloo", "America/Chicago"}
	case "ALP":
		return Location{"Aleppo", "Asia/Damascus"}
	case "ALQ":
		return Location{"Alegrete", "America/Sao_Paulo"}
	case "ALS":
		return Location{"Alamosa", "America/Denver"}
	case "ALW":
		return Location{"Walla Walla", "America/Los_Angeles"}
	case "AMA":
		return Location{"Amarillo", "America/Chicago"}
	case "AMD":
		return Location{"Ahmedabad", "Asia/Kolkata"}
	case "AMH":
		return Location{"", "Africa/Addis_Ababa"}
	case "AMM":
		return Location{"Amman", "Asia/Amman"}
	case "AMQ":
		return Location{"Ambon", "Asia/Jayapura"}
	case "AMS":
		return Location{"Amsterdam", "Europe/Amsterdam"}
	case "ANC":
		return Location{"Anchorage", "America/Anchorage"}
	case "ANF":
		return Location{"Antofagasta", "America/Santiago"}
	case "ANI":
		return Location{"Aniak", "America/Anchorage"}
	case "ANR":
		return Location{"Antwerp", "Europe/Brussels"}
	case "ANU":
		return Location{"St. George", "America/Antigua"}
	case "ANV":
		return Location{"Anvik", "America/Anchorage"}
	case "ANX":
		return Location{"Andenes", "Europe/Oslo"}
	case "AOE":
		return Location{"Eskisehir", "Europe/Istanbul"}
	case "AOG":
		return Location{"Anshan", "Asia/Shanghai"}
	case "AOI":
		return Location{"Ancona", "Europe/Rome"}
	case "AOJ":
		return Location{"Aomori", "Asia/Tokyo"}
	case "AOK":
		return Location{"Karpathos Island", "Europe/Athens"}
	case "AOO":
		return Location{"Altoona", "America/New_York"}
	case "AOR":
		return Location{"Alor Satar", "Asia/Kuala_Lumpur"}
	case "APK":
		return Location{"Apataki", "Pacific/Tahiti"}
	case "APL":
		return Location{"Nampula", "Africa/Maputo"}
	case "APN":
		return Location{"Alpena", "America/Detroit"}
	case "APO":
		return Location{"Carepa", "America/Bogota"}
	case "APW":
		return Location{"Apia", "Pacific/Apia"}
	case "AQG":
		return Location{"Anqing", "Asia/Shanghai"}
	case "AQI":
		return Location{"Qaisumah", "Asia/Riyadh"}
	case "AQJ":
		return Location{"Aqaba", "Asia/Amman"}
	case "AQP":
		return Location{"Arequipa", "America/Lima"}
	case "ARC":
		return Location{"Arctic Village", "America/Anchorage"}
	case "ARD":
		return Location{"Alor Island", "Asia/Makassar"}
	case "ARH":
		return Location{"Archangelsk", "Europe/Moscow"}
	case "ARI":
		return Location{"Arica", "America/Lima"}
	case "ARK":
		return Location{"Arusha", "Africa/Dar_es_Salaam"}
	case "ARM":
		return Location{"Armidale", "Australia/Sydney"}
	case "ARN":
		return Location{"Stockholm", "Europe/Stockholm"}
	case "ART":
		return Location{"Watertown", "America/New_York"}
	case "ARU":
		return Location{"Aracatuba", "America/Sao_Paulo"}
	case "ASB":
		return Location{"Ashgabat", "Asia/Ashgabat"}
	case "ASE":
		return Location{"Aspen", "America/Denver"}
	case "ASF":
		return Location{"Astrakhan", "Europe/Astrakhan"}
	case "ASJ":
		return Location{"Amami", "Asia/Tokyo"}
	case "ASM":
		return Location{"Asmara", "Africa/Asmara"}
	case "ASO":
		return Location{"Asosa", "Africa/Addis_Ababa"}
	case "ASP":
		return Location{"Alice Springs", "Australia/Darwin"}
	case "ASR":
		return Location{"Kayseri", "Europe/Istanbul"}
	case "ASU":
		return Location{"Asuncion", "America/Asuncion"}
	case "ASV":
		return Location{"Amboseli National Park", "Africa/Nairobi"}
	case "ASW":
		return Location{"Aswan", "Africa/Cairo"}
	case "ATA":
		return Location{"Anta", "America/Lima"}
	case "ATD":
		return Location{"Atoifi", "Pacific/Guadalcanal"}
	case "ATH":
		return Location{"Athens", "Europe/Athens"}
	case "ATK":
		return Location{"Atqasuk", "America/Anchorage"}
	case "ATL":
		return Location{"Atlanta", "America/New_York"}
	case "ATM":
		return Location{"Altamira", "America/Santarem"}
	case "ATQ":
		return Location{"Amritsar", "Asia/Kolkata"}
	case "ATW":
		return Location{"Appleton", "America/Chicago"}
	case "ATY":
		return Location{"Watertown", "America/Chicago"}
	case "ATZ":
		return Location{"Assiut", "Africa/Cairo"}
	case "AUA":
		return Location{"Oranjestad", "America/Aruba"}
	case "AUC":
		return Location{"Arauca", "America/Bogota"}
	case "AUG":
		return Location{"Augusta", "America/New_York"}
	case "AUH":
		return Location{"Abu Dhabi", "Asia/Dubai"}
	case "AUK":
		return Location{"Alakanuk", "America/Nome"}
	case "AUQ":
		return Location{"", "Pacific/Marquesas"}
	case "AUR":
		return Location{"Aurillac", "Europe/Paris"}
	case "AUS":
		return Location{"Austin", "America/Chicago"}
	case "AUU":
		return Location{"", "Australia/Brisbane"}
	case "AUX":
		return Location{"Araguaina", "America/Araguaina"}
	case "AUY":
		return Location{"Anelghowhat", "Pacific/Efate"}
	case "AVA":
		return Location{"Anshun", "Asia/Shanghai"}
	case "AVL":
		return Location{"Asheville", "America/New_York"}
	case "AVP":
		return Location{"Wilkes-Barre/Scranton", "America/New_York"}
	case "AVV":
		return Location{"Melbourne", "Australia/Melbourne"}
	case "AWA":
		return Location{"Awassa", "Africa/Addis_Ababa"}
	case "AWD":
		return Location{"Aniwa", "Pacific/Efate"}
	case "AWZ":
		return Location{"Ahwaz", "Asia/Tehran"}
	case "AXA":
		return Location{"The Valley", "America/Anguilla"}
	case "AXD":
		return Location{"Alexandroupolis", "Europe/Athens"}
	case "AXJ":
		return Location{"", "Asia/Tokyo"}
	case "AXM":
		return Location{"Armenia", "America/Bogota"}
	case "AXP":
		return Location{"Spring Point", "America/Nassau"}
	case "AXR":
		return Location{"", "Pacific/Tahiti"}
	case "AXT":
		return Location{"Akita", "Asia/Tokyo"}
	case "AYP":
		return Location{"Ayacucho", "America/Lima"}
	case "AYQ":
		return Location{"Ayers Rock", "Australia/Darwin"}
	case "AYT":
		return Location{"Antalya", "Europe/Istanbul"}
	case "AZA":
		return Location{"Phoenix", "America/Phoenix"}
	case "AZD":
		return Location{"Yazd", "Asia/Tehran"}
	case "AZI":
		return Location{"", "Asia/Dubai"}
	case "AZO":
		return Location{"Kalamazoo", "America/Detroit"}
	case "AZR":
		return Location{"", "Africa/Algiers"}
	case "AZS":
		return Location{"Samana", "America/Santo_Domingo"}
	case "BAG":
		return Location{"Baguio City", "Asia/Manila"}
	case "BAH":
		return Location{"Manama", "Asia/Bahrain"}
	case "BAL":
		return Location{"Batman", "Europe/Istanbul"}
	case "BAQ":
		return Location{"Barranquilla", "America/Bogota"}
	case "BAS":
		return Location{"Ballalae", "Pacific/Guadalcanal"}
	case "BAU":
		return Location{"Bauru", "America/Sao_Paulo"}
	case "BAV":
		return Location{"Baotou", "Asia/Shanghai"}
	case "BAX":
		return Location{"Barnaul", "Asia/Barnaul"}
	case "BAY":
		return Location{"Baia Mare", "Europe/Bucharest"}
	case "BAZ":
		return Location{"Barcelos", "America/Manaus"}
	case "BBA":
		return Location{"Balmaceda", "America/Santiago"}
	case "BBI":
		return Location{"Bhubaneswar", "Asia/Kolkata"}
	case "BBK":
		return Location{"Kasane", "Africa/Gaborone"}
	case "BBN":
		return Location{"Bario", "Asia/Kuching"}
	case "BBU":
		return Location{"Bucharest", "Europe/Bucharest"}
	case "BCD":
		return Location{"Bacolod City", "Asia/Manila"}
	case "BCI":
		return Location{"Barcaldine", "Australia/Brisbane"}
	case "BCM":
		return Location{"Bacau", "Europe/Bucharest"}
	case "BCN":
		return Location{"Barcelona", "Europe/Madrid"}
	case "BCO":
		return Location{"Baco", "Africa/Addis_Ababa"}
	case "BCT":
		return Location{"Boca Raton", "America/New_York"}
	case "BDA":
		return Location{"Hamilton", "Atlantic/Bermuda"}
	case "BDB":
		return Location{"Bundaberg", "Australia/Brisbane"}
	case "BDD":
		return Location{"", "Australia/Brisbane"}
	case "BDJ":
		return Location{"Banjarmasin-Borneo Island", "Asia/Makassar"}
	case "BDL":
		return Location{"Hartford", "America/New_York"}
	case "BDO":
		return Location{"Bandung-Java Island", "Asia/Jakarta"}
	case "BDP":
		return Location{"Bhadrapur", "Asia/Kathmandu"}
	case "BDQ":
		return Location{"Vadodara", "Asia/Kolkata"}
	case "BDR":
		return Location{"Bridgeport", "America/New_York"}
	case "BDS":
		return Location{"Brindisi", "Europe/Rome"}
	case "BDU":
		return Location{"Malselv", "Europe/Oslo"}
	case "BEB":
		return Location{"Balivanich", "Europe/London"}
	case "BEG":
		return Location{"Belgrad", "Europe/Belgrade"}
	case "BEJ":
		return Location{"Tanjung Redep-Borneo Island", "Asia/Makassar"}
	case "BEK":
		return Location{"", "Asia/Kolkata"}
	case "BEL":
		return Location{"Belem", "America/Belem"}
	case "BEM":
		return Location{"Bossembele", "Africa/Bangui"}
	case "BEN":
		return Location{"Benghazi", "Africa/Tripoli"}
	case "BER":
		return Location{"Berlin", "Europe/Berlin"}
	case "BES":
		return Location{"Brest/Guipavas", "Europe/Paris"}
	case "BET":
		return Location{"Bethel", "America/Anchorage"}
	case "BEU":
		return Location{"", "Australia/Brisbane"}
	case "BEW":
		return Location{"Beira", "Africa/Maputo"}
	case "BEY":
		return Location{"Beirut", "Asia/Beirut"}
	case "BFD":
		return Location{"Bradford", "America/New_York"}
	case "BFF":
		return Location{"Scottsbluff", "America/Denver"}
	case "BFI":
		return Location{"Seattle", "America/Los_Angeles"}
	case "BFJ":
		return Location{"", "Pacific/Fiji"}
	case "BFL":
		return Location{"Bakersfield", "America/Los_Angeles"}
	case "BFM":
		return Location{"Mobile", "America/Chicago"}
	case "BFN":
		return Location{"Bloemfontain", "Africa/Johannesburg"}
	case "BFS":
		return Location{"Belfast", "Europe/London"}
	case "BFV":
		return Location{"", "Asia/Bangkok"}
	case "BFX":
		return Location{"Bafoussam", "Africa/Douala"}
	case "BGA":
		return Location{"Bucaramanga", "America/Bogota"}
	case "BGF":
		return Location{"Bangui", "Africa/Bangui"}
	case "BGI":
		return Location{"Bridgetown", "America/Barbados"}
	case "BGM":
		return Location{"Binghamton", "America/New_York"}
	case "BGN":
		return Location{"Belaya Gora", "Asia/Magadan"}
	case "BGO":
		return Location{"Bergen", "Europe/Oslo"}
	case "BGR":
		return Location{"Bangor", "America/New_York"}
	case "BGS":
		return Location{"Big Spring", "America/Chicago"}
	case "BGW":
		return Location{"Baghdad", "Asia/Baghdad"}
	case "BGX":
		return Location{"Bage", "America/Sao_Paulo"}
	case "BGY":
		return Location{"Bergamo", "Europe/Rome"}
	case "BHB":
		return Location{"Bar Harbor", "America/New_York"}
	case "BHD":
		return Location{"Belfast", "Europe/London"}
	case "BHE":
		return Location{"Blenheim", "Pacific/Auckland"}
	case "BHH":
		return Location{"", "Asia/Riyadh"}
	case "BHI":
		return Location{"Bahia Blanca", "America/Argentina/Buenos_Aires"}
	case "BHJ":
		return Location{"Bhuj", "Asia/Kolkata"}
	case "BHK":
		return Location{"Bukhara", "Asia/Samarkand"}
	case "BHM":
		return Location{"Birmingham", "America/Chicago"}
	case "BHO":
		return Location{"Bhopal", "Asia/Kolkata"}
	case "BHQ":
		return Location{"Broken Hill", "Australia/Broken_Hill"}
	case "BHR":
		return Location{"Bharatpur", "Asia/Kathmandu"}
	case "BHU":
		return Location{"Bhavnagar", "Asia/Kolkata"}
	case "BHX":
		return Location{"Birmingham", "Europe/London"}
	case "BHY":
		return Location{"Beihai", "Asia/Shanghai"}
	case "BIA":
		return Location{"Bastia/Poretta", "Europe/Paris"}
	case "BIH":
		return Location{"Bishop", "America/Los_Angeles"}
	case "BIK":
		return Location{"Biak-Supiori Island", "Asia/Jayapura"}
	case "BIL":
		return Location{"Billings", "America/Denver"}
	case "BIM":
		return Location{"South Bimini", "America/Nassau"}
	case "BIO":
		return Location{"Bilbao", "Europe/Madrid"}
	case "BIQ":
		return Location{"Biarritz/Anglet/Bayonne", "Europe/Paris"}
	case "BIR":
		return Location{"Biratnagar", "Asia/Kathmandu"}
	case "BIS":
		return Location{"Bismarck", "America/Chicago"}
	case "BJA":
		return Location{"Bejaia", "Africa/Algiers"}
	case "BJC":
		return Location{"Denver", "America/Denver"}
	case "BJF":
		return Location{"Batsfjord", "Europe/Oslo"}
	case "BJI":
		return Location{"Bemidji", "America/Chicago"}
	case "BJL":
		return Location{"Banjul", "Africa/Banjul"}
	case "BJM":
		return Location{"Bujumbura", "Africa/Bujumbura"}
	case "BJR":
		return Location{"Bahir Dar", "Africa/Addis_Ababa"}
	case "BJV":
		return Location{"Bodrum", "Europe/Istanbul"}
	case "BJW":
		return Location{"Bajawa", "Asia/Makassar"}
	case "BJX":
		return Location{"Silao", "America/Mexico_City"}
	case "BJZ":
		return Location{"Badajoz", "Europe/Madrid"}
	case "BKA":
		return Location{"Moscow", "Europe/Moscow"}
	case "BKB":
		return Location{"Bikaner", "Asia/Kolkata"}
	case "BKC":
		return Location{"Buckland", "America/Anchorage"}
	case "BKG":
		return Location{"Branson", "America/Chicago"}
	case "BKI":
		return Location{"Kota Kinabalu", "Asia/Kuching"}
	case "BKK":
		return Location{"Bangkok", "Asia/Bangkok"}
	case "BKM":
		return Location{"Bakalalan", "Asia/Kuching"}
	case "BKO":
		return Location{"Senou", "Africa/Bamako"}
	case "BKQ":
		return Location{"Blackall", "Australia/Brisbane"}
	case "BKS":
		return Location{"Bengkulu-Sumatra Island", "Asia/Jakarta"}
	case "BKW":
		return Location{"Beckley", "America/New_York"}
	case "BKZ":
		return Location{"Bukoba", "Africa/Dar_es_Salaam"}
	case "BLA":
		return Location{"Barcelona", "America/Caracas"}
	case "BLB":
		return Location{"Panama City", "America/Panama"}
	case "BLD":
		return Location{"Boulder City", "America/Los_Angeles"}
	case "BLI":
		return Location{"Bellingham", "America/Los_Angeles"}
	case "BLJ":
		return Location{"Batna", "Africa/Algiers"}
	case "BLL":
		return Location{"Billund", "Europe/Copenhagen"}
	case "BLQ":
		return Location{"Bologna", "Europe/Rome"}
	case "BLR":
		return Location{"Bangalore", "Asia/Kolkata"}
	case "BLV":
		return Location{"Belleville", "America/Chicago"}
	case "BLZ":
		return Location{"Blantyre", "Africa/Blantyre"}
	case "BMA":
		return Location{"Stockholm", "Europe/Stockholm"}
	case "BME":
		return Location{"Broome", "Australia/Perth"}
	case "BMI":
		return Location{"Bloomington-Normal", "America/Chicago"}
	case "BMO":
		return Location{"Banmaw", "Asia/Yangon"}
	case "BMU":
		return Location{"Bima-Sumbawa Island", "Asia/Makassar"}
	case "BMV":
		return Location{"Buon Ma Thuot", "Asia/Ho_Chi_Minh"}
	case "BMW":
		return Location{"Bordj Badji Mokhtar", "Africa/Algiers"}
	case "BNA":
		return Location{"Nashville", "America/Chicago"}
	case "BND":
		return Location{"Bandar Abbas", "Asia/Tehran"}
	case "BNE":
		return Location{"Brisbane", "Australia/Brisbane"}
	case "BNI":
		return Location{"Benin", "Africa/Lagos"}
	case "BNK":
		return Location{"Ballina", "Australia/Sydney"}
	case "BNN":
		return Location{"Bronnoy", "Europe/Oslo"}
	case "BNS":
		return Location{"Barinas", "America/Caracas"}
	case "BNX":
		return Location{"Banja Luka", "Europe/Sarajevo"}
	case "BNY":
		return Location{"Anua", "Pacific/Guadalcanal"}
	case "BOB":
		return Location{"Motu Mute", "Pacific/Tahiti"}
	case "BOC":
		return Location{"Isla Colon", "America/Panama"}
	case "BOD":
		return Location{"Bordeaux/Merignac", "Europe/Paris"}
	case "BOG":
		return Location{"Bogota", "America/Bogota"}
	case "BOH":
		return Location{"Bournemouth", "Europe/London"}
	case "BOI":
		return Location{"Boise", "America/Boise"}
	case "BOJ":
		return Location{"Burgas", "Europe/Sofia"}
	case "BOM":
		return Location{"Mumbai", "Asia/Kolkata"}
	case "BON":
		return Location{"Kralendijk", "America/Kralendijk"}
	case "BOO":
		return Location{"Bodo", "Europe/Oslo"}
	case "BOS":
		return Location{"Boston", "America/New_York"}
	case "BOY":
		return Location{"Bobo Dioulasso", "Africa/Ouagadougou"}
	case "BPG":
		return Location{"Barra Do Garcas", "America/Cuiaba"}
	case "BPL":
		return Location{"Bole", "Asia/Shanghai"}
	case "BPN":
		return Location{"Balikpapan-Borneo Island", "Asia/Makassar"}
	case "BPS":
		return Location{"Porto Seguro", "America/Bahia"}
	case "BPT":
		return Location{"Beaumont/Port Arthur", "America/Chicago"}
	case "BPX":
		return Location{"Bangda", "Asia/Shanghai"}
	case "BQB":
		return Location{"Busselton", "Australia/Perth"}
	case "BQD":
		return Location{"Budardalur", "Atlantic/Reykjavik"}
	case "BQK":
		return Location{"Brunswick", "America/New_York"}
	case "BQL":
		return Location{"", "Australia/Brisbane"}
	case "BQN":
		return Location{"Aguadilla", "America/Puerto_Rico"}
	case "BQS":
		return Location{"Blagoveschensk", "Asia/Yakutsk"}
	case "BQW":
		return Location{"", "Australia/Perth"}
	case "BRA":
		return Location{"Barreiras", "America/Bahia"}
	case "BRB":
		return Location{"", "America/Fortaleza"}
	case "BRC":
		return Location{"San Carlos de Bariloche", "America/Argentina/Salta"}
	case "BRD":
		return Location{"Brainerd", "America/Chicago"}
	case "BRE":
		return Location{"Bremen", "Europe/Berlin"}
	case "BRI":
		return Location{"Bari", "Europe/Rome"}
	case "BRK":
		return Location{"", "Australia/Sydney"}
	case "BRL":
		return Location{"Burlington", "America/Chicago"}
	case "BRM":
		return Location{"Barquisimeto", "America/Caracas"}
	case "BRN":
		return Location{"Bern", "Europe/Zurich"}
	case "BRO":
		return Location{"Brownsville", "America/Chicago"}
	case "BRQ":
		return Location{"Brno", "Europe/Prague"}
	case "BRR":
		return Location{"Eoligarry", "Europe/London"}
	case "BRS":
		return Location{"Bristol", "Europe/London"}
	case "BRU":
		return Location{"Brussels", "Europe/Brussels"}
	case "BRW":
		return Location{"Barrow", "America/Anchorage"}
	case "BSA":
		return Location{"Bosaso", "Africa/Mogadishu"}
	case "BSB":
		return Location{"Brasilia", "America/Sao_Paulo"}
	case "BSC":
		return Location{"Bahia Solano", "America/Bogota"}
	case "BSD":
		return Location{"", "Asia/Shanghai"}
	case "BSG":
		return Location{"", "Africa/Malabo"}
	case "BSK":
		return Location{"Biskra", "Africa/Algiers"}
	case "BSL":
		return Location{"Bale/Mulhouse", "Europe/Paris"}
	case "BSO":
		return Location{"Basco", "Asia/Manila"}
	case "BSR":
		return Location{"Basrah", "Asia/Baghdad"}
	case "BTC":
		return Location{"Batticaloa", "Asia/Colombo"}
	case "BTH":
		return Location{"Batam Island", "Asia/Jakarta"}
	case "BTI":
		return Location{"Barter Island Lrrs", "America/Anchorage"}
	case "BTJ":
		return Location{"Banda Aceh-Sumatra Island", "Asia/Jakarta"}
	case "BTK":
		return Location{"Bratsk", "Asia/Irkutsk"}
	case "BTM":
		return Location{"Butte", "America/Denver"}
	case "BTR":
		return Location{"Baton Rouge", "America/Chicago"}
	case "BTS":
		return Location{"Bratislava", "Europe/Bratislava"}
	case "BTT":
		return Location{"Bettles", "America/Anchorage"}
	case "BTU":
		return Location{"Bintulu", "Asia/Kuching"}
	case "BTV":
		return Location{"Burlington", "America/New_York"}
	case "BTW":
		return Location{"Batu Licin-Borneo Island", "Asia/Makassar"}
	case "BUA":
		return Location{"Buka Island", "Pacific/Bougainville"}
	case "BUC":
		return Location{"", "Australia/Brisbane"}
	case "BUD":
		return Location{"Budapest", "Europe/Budapest"}
	case "BUF":
		return Location{"Buffalo", "America/New_York"}
	case "BUN":
		return Location{"Buenaventura", "America/Bogota"}
	case "BUP":
		return Location{"", "Asia/Kolkata"}
	case "BUQ":
		return Location{"Bulawayo", "Africa/Harare"}
	case "BUR":
		return Location{"Burbank", "America/Los_Angeles"}
	case "BUS":
		return Location{"Batumi", "Asia/Tbilisi"}
	case "BUW":
		return Location{"Bau Bau-Butung Island", "Asia/Makassar"}
	case "BUX":
		return Location{"", "Africa/Lubumbashi"}
	case "BUZ":
		return Location{"Bushehr", "Asia/Tehran"}
	case "BVA":
		return Location{"Beauvais/Tille", "Europe/Paris"}
	case "BVB":
		return Location{"Boa Vista", "America/Boa_Vista"}
	case "BVC":
		return Location{"Rabil", "Atlantic/Cape_Verde"}
	case "BVE":
		return Location{"Brive-la-Gaillarde", "Europe/Paris"}
	case "BVG":
		return Location{"Berlevag", "Europe/Oslo"}
	case "BVH":
		return Location{"Vilhena", "America/Cuiaba"}
	case "BVI":
		return Location{"", "Australia/Brisbane"}
	case "BVS":
		return Location{"Breves", "America/Belem"}
	case "BWA":
		return Location{"Bhairawa", "Asia/Kathmandu"}
	case "BWI":
		return Location{"Baltimore", "America/New_York"}
	case "BWK":
		return Location{"Brac Island", "Europe/Zagreb"}
	case "BWN":
		return Location{"Bandar Seri Begawan", "Asia/Brunei"}
	case "BWT":
		return Location{"Burnie", "Australia/Hobart"}
	case "BXG":
		return Location{"", "Australia/Melbourne"}
	case "BXU":
		return Location{"Butuan City", "Asia/Manila"}
	case "BXY":
		return Location{"Baikonur", "Asia/Qyzylorda"}
	case "BYC":
		return Location{"Yacuiba", "America/La_Paz"}
	case "BYK":
		return Location{"", "Africa/Abidjan"}
	case "BYN":
		return Location{"Bayankhongor", "Asia/Ulaanbaatar"}
	case "BYO":
		return Location{"Bonito", "America/Campo_Grande"}
	case "BZE":
		return Location{"Belize City", "America/Belize"}
	case "BZG":
		return Location{"Bydgoszcz", "Europe/Warsaw"}
	case "BZL":
		return Location{"Barisal", "Asia/Dhaka"}
	case "BZN":
		return Location{"Bozeman", "America/Denver"}
	case "BZO":
		return Location{"Bolzano", "Europe/Rome"}
	case "BZR":
		return Location{"Beziers/Vias", "Europe/Paris"}
	case "BZV":
		return Location{"Brazzaville", "Africa/Brazzaville"}
	case "CAB":
		return Location{"Cabinda", "Africa/Luanda"}
	case "CAC":
		return Location{"Cascavel", "America/Sao_Paulo"}
	case "CAE":
		return Location{"Columbia", "America/New_York"}
	case "CAG":
		return Location{"Cagliari", "Europe/Rome"}
	case "CAH":
		return Location{"Ca Mau City", "Asia/Ho_Chi_Minh"}
	case "CAI":
		return Location{"Cairo", "Africa/Cairo"}
	case "CAK":
		return Location{"Akron", "America/New_York"}
	case "CAL":
		return Location{"Campbeltown", "Europe/London"}
	case "CAN":
		return Location{"Guangzhou", "Asia/Shanghai"}
	case "CAP":
		return Location{"Cap Haitien", "America/Port-au-Prince"}
	case "CAU":
		return Location{"Caruaru", "America/Recife"}
	case "CAW":
		return Location{"Campos Dos Goytacazes", "America/Sao_Paulo"}
	case "CAY":
		return Location{"Cayenne / Rochambeau", "America/Cayenne"}
	case "CAZ":
		return Location{"", "Australia/Sydney"}
	case "CBB":
		return Location{"Cochabamba", "America/La_Paz"}
	case "CBG":
		return Location{"Cambridge", "Europe/London"}
	case "CBH":
		return Location{"Bechar", "Africa/Algiers"}
	case "CBL":
		return Location{"", "America/Caracas"}
	case "CBO":
		return Location{"Cotabato City", "Asia/Manila"}
	case "CBQ":
		return Location{"Calabar", "Africa/Lagos"}
	case "CBR":
		return Location{"Canberra", "Australia/Sydney"}
	case "CBT":
		return Location{"Catumbela", "Africa/Luanda"}
	case "CCC":
		return Location{"Cayo Coco", "America/Havana"}
	case "CCF":
		return Location{"Carcassonne/Salvaza", "Europe/Paris"}
	case "CCJ":
		return Location{"Calicut", "Asia/Kolkata"}
	case "CCK":
		return Location{"Cocos (Keeling) Islands", "Indian/Cocos"}
	case "CCP":
		return Location{"Concepcion", "America/Santiago"}
	case "CCR":
		return Location{"Concord", "America/Los_Angeles"}
	case "CCS":
		return Location{"Caracas", "America/Caracas"}
	case "CCU":
		return Location{"Kolkata", "Asia/Kolkata"}
	case "CCV":
		return Location{"Craig Cove", "Pacific/Efate"}
	case "CDB":
		return Location{"Cold Bay", "America/Nome"}
	case "CDC":
		return Location{"Cedar City", "America/Denver"}
	case "CDG":
		return Location{"Paris", "Europe/Paris"}
	case "CDP":
		return Location{"", "Asia/Kolkata"}
	case "CDR":
		return Location{"Chadron", "America/Denver"}
	case "CDT":
		return Location{"Calamocha", "Europe/Madrid"}
	case "CDV":
		return Location{"Cordova", "America/Anchorage"}
	case "CEB":
		return Location{"Lapu-Lapu City", "Asia/Manila"}
	case "CEC":
		return Location{"Crescent City", "America/Los_Angeles"}
	case "CED":
		return Location{"", "Australia/Adelaide"}
	case "CEE":
		return Location{"Cherepovets", "Europe/Moscow"}
	case "CEI":
		return Location{"Chiang Rai", "Asia/Bangkok"}
	case "CEK":
		return Location{"Chelyabinsk", "Asia/Yekaterinburg"}
	case "CEM":
		return Location{"Central", "America/Anchorage"}
	case "CEN":
		return Location{"Ciudad Obregon", "America/Hermosillo"}
	case "CEZ":
		return Location{"Cortez", "America/Denver"}
	case "CFB":
		return Location{"Cabo Frio", "America/Sao_Paulo"}
	case "CFE":
		return Location{"Clermont-Ferrand/Auvergne", "Europe/Paris"}
	case "CFG":
		return Location{"Cienfuegos", "America/Havana"}
	case "CFN":
		return Location{"Donegal", "Europe/Dublin"}
	case "CFR":
		return Location{"Caen/Carpiquet", "Europe/Paris"}
	case "CFS":
		return Location{"Coffs Harbour", "Australia/Sydney"}
	case "CFU":
		return Location{"Kerkyra Island", "Europe/Athens"}
	case "CGB":
		return Location{"Cuiaba", "America/Cuiaba"}
	case "CGD":
		return Location{"Changde", "Asia/Shanghai"}
	case "CGH":
		return Location{"Sao Paulo", "America/Sao_Paulo"}
	case "CGI":
		return Location{"Cape Girardeau", "America/Chicago"}
	case "CGK":
		return Location{"Jakarta", "Asia/Jakarta"}
	case "CGM":
		return Location{"", "Asia/Manila"}
	case "CGN":
		return Location{"Cologne", "Europe/Berlin"}
	case "CGO":
		return Location{"Zhengzhou", "Asia/Shanghai"}
	case "CGP":
		return Location{"Chittagong", "Asia/Dhaka"}
	case "CGQ":
		return Location{"Changchun", "Asia/Shanghai"}
	case "CGR":
		return Location{"Campo Grande", "America/Campo_Grande"}
	case "CGY":
		return Location{"Cagayan De Oro City", "Asia/Manila"}
	case "CHA":
		return Location{"Chattanooga", "America/New_York"}
	case "CHC":
		return Location{"Christchurch", "Pacific/Auckland"}
	case "CHG":
		return Location{"Chaoyang", "Asia/Shanghai"}
	case "CHH":
		return Location{"Chachapoyas", "America/Lima"}
	case "CHO":
		return Location{"Charlottesville", "America/New_York"}
	case "CHQ":
		return Location{"Souda", "Europe/Athens"}
	case "CHS":
		return Location{"Charleston", "America/New_York"}
	case "CHT":
		return Location{"Waitangi", "Pacific/Chatham"}
	case "CHU":
		return Location{"Chuathbaluk", "America/Anchorage"}
	case "CHX":
		return Location{"Changuinola", "America/Panama"}
	case "CHY":
		return Location{"", "Pacific/Guadalcanal"}
	case "CIA":
		return Location{"Roma", "Europe/Rome"}
	case "CID":
		return Location{"Cedar Rapids", "America/Chicago"}
	case "CIF":
		return Location{"Chifeng", "Asia/Shanghai"}
	case "CIH":
		return Location{"Changzhi", "Asia/Shanghai"}
	case "CIJ":
		return Location{"Cobija", "America/La_Paz"}
	case "CIK":
		return Location{"Chalkyitsik", "America/Anchorage"}
	case "CIT":
		return Location{"Shymkent", "Asia/Almaty"}
	case "CIU":
		return Location{"Sault Ste Marie", "America/Detroit"}
	case "CIX":
		return Location{"Chiclayo", "America/Lima"}
	case "CIY":
		return Location{"Comiso", "Europe/Rome"}
	case "CIZ":
		return Location{"Coari", "America/Manaus"}
	case "CJA":
		return Location{"Cajamarca", "America/Lima"}
	case "CJB":
		return Location{"Coimbatore", "Asia/Kolkata"}
	case "CJC":
		return Location{"Calama", "America/Santiago"}
	case "CJJ":
		return Location{"Cheongju", "Asia/Seoul"}
	case "CJM":
		return Location{"", "Asia/Bangkok"}
	case "CJS":
		return Location{"Ciudad Juarez", "America/Ojinaga"}
	case "CJU":
		return Location{"Jeju City", "Asia/Seoul"}
	case "CKB":
		return Location{"Clarksburg", "America/New_York"}
	case "CKG":
		return Location{"Chongqing", "Asia/Shanghai"}
	case "CKH":
		return Location{"Chokurdah", "Asia/Srednekolymsk"}
	case "CKO":
		return Location{"Cornelio Procopio", "America/Sao_Paulo"}
	case "CKS":
		return Location{"Carajas", "America/Belem"}
	case "CKY":
		return Location{"Conakry", "Africa/Conakry"}
	case "CKZ":
		return Location{"Canakkale", "Europe/Istanbul"}
	case "CLD":
		return Location{"Carlsbad", "America/Los_Angeles"}
	case "CLE":
		return Location{"Cleveland", "America/New_York"}
	case "CLJ":
		return Location{"Cluj-Napoca", "Europe/Bucharest"}
	case "CLL":
		return Location{"College Station", "America/Chicago"}
	case "CLO":
		return Location{"Cali", "America/Bogota"}
	case "CLP":
		return Location{"Clarks Point", "America/Anchorage"}
	case "CLQ":
		return Location{"Colima", "America/Mexico_City"}
	case "CLT":
		return Location{"Charlotte", "America/New_York"}
	case "CLV":
		return Location{"Caldas Novas", "America/Sao_Paulo"}
	case "CLY":
		return Location{"Calvi/Sainte-Catherine", "Europe/Paris"}
	case "CMA":
		return Location{"", "Australia/Brisbane"}
	case "CMB":
		return Location{"Colombo", "Asia/Colombo"}
	case "CME":
		return Location{"Ciudad del Carmen", "America/Merida"}
	case "CMF":
		return Location{"Chambery/Aix-les-Bains", "Europe/Paris"}
	case "CMG":
		return Location{"Corumba", "America/Campo_Grande"}
	case "CMH":
		return Location{"Columbus", "America/New_York"}
	case "CMI":
		return Location{"Champaign/Urbana", "America/Chicago"}
	case "CMN":
		return Location{"Casablanca", "Africa/Casablanca"}
	case "CMW":
		return Location{"Camaguey", "America/Havana"}
	case "CMX":
		return Location{"Hancock", "America/Detroit"}
	case "CNC":
		return Location{"", "Australia/Brisbane"}
	case "CND":
		return Location{"Constanta", "Europe/Bucharest"}
	case "CNF":
		return Location{"Belo Horizonte", "America/Sao_Paulo"}
	case "CNJ":
		return Location{"Cloncurry", "Australia/Brisbane"}
	case "CNM":
		return Location{"Carlsbad", "America/Denver"}
	case "CNN":
		return Location{"Chulman", "Asia/Yakutsk"}
	case "CNQ":
		return Location{"Corrientes", "America/Argentina/Cordoba"}
	case "CNS":
		return Location{"Cairns", "Australia/Brisbane"}
	case "CNX":
		return Location{"Chiang Mai", "Asia/Bangkok"}
	case "CNY":
		return Location{"Moab", "America/Denver"}
	case "COD":
		return Location{"Cody", "America/Denver"}
	case "COH":
		return Location{"", "Asia/Kolkata"}
	case "COK":
		return Location{"Cochin", "Asia/Kolkata"}
	case "COO":
		return Location{"Cotonou", "Africa/Porto-Novo"}
	case "COR":
		return Location{"Cordoba", "America/Argentina/Cordoba"}
	case "COS":
		return Location{"Colorado Springs", "America/Denver"}
	case "COU":
		return Location{"Columbia", "America/Chicago"}
	case "CPC":
		return Location{"Chapelco/San Martin de los Andes", "America/Argentina/Salta"}
	case "CPD":
		return Location{"", "Australia/Adelaide"}
	case "CPE":
		return Location{"Campeche", "America/Merida"}
	case "CPH":
		return Location{"Copenhagen", "Europe/Copenhagen"}
	case "CPO":
		return Location{"Copiapo", "America/Santiago"}
	case "CPR":
		return Location{"Casper", "America/Denver"}
	case "CPT":
		return Location{"Cape Town", "Africa/Johannesburg"}
	case "CPV":
		return Location{"Campina Grande", "America/Fortaleza"}
	case "CPX":
		return Location{"Culebra Island", "America/Puerto_Rico"}
	case "CQD":
		return Location{"Shahrekord", "Asia/Tehran"}
	case "CRA":
		return Location{"Craiova", "Europe/Bucharest"}
	case "CRD":
		return Location{"Comodoro Rivadavia", "America/Argentina/Catamarca"}
	case "CRI":
		return Location{"Colonel Hill", "America/Nassau"}
	case "CRK":
		return Location{"Angeles City", "Asia/Manila"}
	case "CRL":
		return Location{"Brussels", "Europe/Brussels"}
	case "CRM":
		return Location{"Catarman", "Asia/Manila"}
	case "CRP":
		return Location{"Corpus Christi", "America/Chicago"}
	case "CRV":
		return Location{"Crotone", "Europe/Rome"}
	case "CRW":
		return Location{"Charleston", "America/New_York"}
	case "CRZ":
		return Location{"Turkmenabat", "Asia/Ashgabat"}
	case "CSG":
		return Location{"Columbus", "America/New_York"}
	case "CSK":
		return Location{"Cap Skirring", "Africa/Dakar"}
	case "CSU":
		return Location{"Santa Cruz Do Sul", "America/Sao_Paulo"}
	case "CSX":
		return Location{"Changsha", "Asia/Shanghai"}
	case "CSY":
		return Location{"Cheboksary", "Europe/Moscow"}
	case "CTA":
		return Location{"Catania", "Europe/Rome"}
	case "CTC":
		return Location{"Catamarca", "America/Argentina/Catamarca"}
	case "CTD":
		return Location{"Chitre", "America/Panama"}
	case "CTG":
		return Location{"Cartagena", "America/Bogota"}
	case "CTL":
		return Location{"Charleville", "Australia/Brisbane"}
	case "CTM":
		return Location{"Chetumal", "America/Cancun"}
	case "CTS":
		return Location{"Chitose / Tomakomai", "Asia/Tokyo"}
	case "CTU":
		return Location{"Chengdu", "Asia/Shanghai"}
	case "CUC":
		return Location{"Cucuta", "America/Bogota"}
	case "CUE":
		return Location{"Cuenca", "America/Guayaquil"}
	case "CUF":
		return Location{"Cuneo", "Europe/Rome"}
	case "CUL":
		return Location{"Culiacan", "America/Mazatlan"}
	case "CUM":
		return Location{"", "America/Caracas"}
	case "CUN":
		return Location{"Cancun", "America/Cancun"}
	case "CUR":
		return Location{"Willemstad", "America/Curacao"}
	case "CUU":
		return Location{"Chihuahua", "America/Chihuahua"}
	case "CUZ":
		return Location{"Cusco", "America/Lima"}
	case "CVG":
		return Location{"Hebron", "America/New_York"}
	case "CVM":
		return Location{"Ciudad Victoria", "America/Monterrey"}
	case "CVN":
		return Location{"Clovis", "America/Denver"}
	case "CVQ":
		return Location{"", "Australia/Perth"}
	case "CVU":
		return Location{"Corvo", "Atlantic/Azores"}
	case "CWA":
		return Location{"Mosinee", "America/Chicago"}
	case "CWB":
		return Location{"Curitiba", "America/Sao_Paulo"}
	case "CWL":
		return Location{"Cardiff", "Europe/London"}
	case "CXB":
		return Location{"Cox's Bazar", "Asia/Dhaka"}
	case "CXI":
		return Location{"Banana", "Pacific/Kiritimati"}
	case "CXJ":
		return Location{"Caxias Do Sul", "America/Sao_Paulo"}
	case "CXR":
		return Location{"Nha Trang", "Asia/Ho_Chi_Minh"}
	case "CXY":
		return Location{"Cat Cay", "America/Nassau"}
	case "CYA":
		return Location{"Les Cayes", "America/Port-au-Prince"}
	case "CYB":
		return Location{"Cayman Brac", "America/Cayman"}
	case "CYF":
		return Location{"Chefornak", "America/Nome"}
	case "CYI":
		return Location{"Chiayi City", "Asia/Taipei"}
	case "CYO":
		return Location{"Cayo Largo del Sur", "America/Havana"}
	case "CYP":
		return Location{"Calbayog City", "Asia/Manila"}
	case "CYS":
		return Location{"Cheyenne", "America/Denver"}
	case "CYX":
		return Location{"Cherskiy", "Asia/Srednekolymsk"}
	case "CYZ":
		return Location{"Cauayan City", "Asia/Manila"}
	case "CZL":
		return Location{"Constantine", "Africa/Algiers"}
	case "CZM":
		return Location{"Cozumel", "America/Cancun"}
	case "CZS":
		return Location{"Cruzeiro Do Sul", "America/Rio_Branco"}
	case "CZU":
		return Location{"Corozal", "America/Bogota"}
	case "CZX":
		return Location{"Changzhou", "Asia/Shanghai"}
	case "DAB":
		return Location{"Daytona Beach", "America/New_York"}
	case "DAC":
		return Location{"Dhaka", "Asia/Dhaka"}
	case "DAD":
		return Location{"Da Nang", "Asia/Ho_Chi_Minh"}
	case "DAL":
		return Location{"Dallas", "America/Chicago"}
	case "DAM":
		return Location{"Damascus", "Asia/Damascus"}
	case "DAR":
		return Location{"Dar es Salaam", "Africa/Dar_es_Salaam"}
	case "DAT":
		return Location{"Datong", "Asia/Shanghai"}
	case "DAU":
		return Location{"Daru", "Pacific/Port_Moresby"}
	case "DAV":
		return Location{"David", "America/Panama"}
	case "DAY":
		return Location{"Dayton", "America/New_York"}
	case "DBO":
		return Location{"Dubbo", "Australia/Sydney"}
	case "DBQ":
		return Location{"Dubuque", "America/Chicago"}
	case "DBR":
		return Location{"Darbhanga", "Asia/Kolkata"}
	case "DBV":
		return Location{"Dubrovnik", "Europe/Zagreb"}
	case "DCA":
		return Location{"Washington", "America/New_York"}
	case "DCM":
		return Location{"Castres/Mazamet", "Europe/Paris"}
	case "DDC":
		return Location{"Dodge City", "America/Chicago"}
	case "DDG":
		return Location{"Dandong", "Asia/Shanghai"}
	case "DEB":
		return Location{"Debrecen", "Europe/Budapest"}
	case "DEC":
		return Location{"Decatur", "America/Chicago"}
	case "DED":
		return Location{"Dehradun", "Asia/Kolkata"}
	case "DEE":
		return Location{"Kunashir Island", "Asia/Ust-Nera"}
	case "DEL":
		return Location{"New Delhi", "Asia/Kolkata"}
	case "DEN":
		return Location{"Denver", "America/Denver"}
	case "DFW":
		return Location{"Dallas-Fort Worth", "America/Chicago"}
	case "DGE":
		return Location{"Mudgee", "Australia/Sydney"}
	case "DGH":
		return Location{"Deoghar", "Asia/Kolkata"}
	case "DGO":
		return Location{"Durango", "America/Monterrey"}
	case "DGT":
		return Location{"Dumaguete City", "Asia/Manila"}
	case "DHI":
		return Location{"Dhangarhi", "Asia/Kathmandu"}
	case "DHM":
		return Location{"", "Asia/Kolkata"}
	case "DHN":
		return Location{"Dothan", "America/Chicago"}
	case "DIB":
		return Location{"Dibrugarh", "Asia/Kolkata"}
	case "DIE":
		return Location{"", "Indian/Antananarivo"}
	case "DIG":
		return Location{"Shangri-La", "Asia/Shanghai"}
	case "DIK":
		return Location{"Dickinson", "America/Denver"}
	case "DIL":
		return Location{"Dili", "Asia/Dili"}
	case "DIN":
		return Location{"Dien Bien Phu", "Asia/Bangkok"}
	case "DIR":
		return Location{"Dire Dawa", "Africa/Addis_Ababa"}
	case "DIU":
		return Location{"Diu", "Asia/Kolkata"}
	case "DIY":
		return Location{"Diyarbakir", "Europe/Istanbul"}
	case "DJB":
		return Location{"Jambi-Sumatra Island", "Asia/Jakarta"}
	case "DJE":
		return Location{"Djerba", "Africa/Tunis"}
	case "DJG":
		return Location{"Djanet", "Africa/Algiers"}
	case "DJJ":
		return Location{"Jayapura-Papua Island", "Asia/Jayapura"}
	case "DKR":
		return Location{"Dakar", "Africa/Dakar"}
	case "DKS":
		return Location{"Dikson", "Asia/Krasnoyarsk"}
	case "DLA":
		return Location{"Douala", "Africa/Douala"}
	case "DLC":
		return Location{"Dalian", "Asia/Shanghai"}
	case "DLE":
		return Location{"Dole/Tavaux", "Europe/Paris"}
	case "DLG":
		return Location{"Dillingham", "America/Anchorage"}
	case "DLH":
		return Location{"Duluth", "America/Chicago"}
	case "DLI":
		return Location{"Dalat", "Asia/Ho_Chi_Minh"}
	case "DLM":
		return Location{"Dalaman", "Europe/Istanbul"}
	case "DLU":
		return Location{"Xiaguan", "Asia/Shanghai"}
	case "DLY":
		return Location{"Dillon's Bay", "Pacific/Efate"}
	case "DLZ":
		return Location{"Dalanzadgad", "Asia/Ulaanbaatar"}
	case "DMB":
		return Location{"Taraz", "Asia/Almaty"}
	case "DMD":
		return Location{"", "Australia/Brisbane"}
	case "DME":
		return Location{"Moscow", "Europe/Moscow"}
	case "DMK":
		return Location{"Bangkok", "Asia/Bangkok"}
	case "DMM":
		return Location{"Ad Dammam", "Asia/Riyadh"}
	case "DMU":
		return Location{"Dimapur", "Asia/Kolkata"}
	case "DND":
		return Location{"Dundee", "Europe/London"}
	case "DNH":
		return Location{"Dunhuang", "Asia/Shanghai"}
	case "DNZ":
		return Location{"Denizli", "Europe/Istanbul"}
	case "DOB":
		return Location{"Dobo-Kobror Island", "Asia/Jayapura"}
	case "DOD":
		return Location{"Dodoma", "Africa/Dar_es_Salaam"}
	case "DOH":
		return Location{"Doha", "Asia/Qatar"}
	case "DOL":
		return Location{"Deauville", "Europe/Paris"}
	case "DOM":
		return Location{"Marigot", "America/Dominica"}
	case "DOY":
		return Location{"Dongying", "Asia/Shanghai"}
	case "DPL":
		return Location{"Dipolog City", "Asia/Manila"}
	case "DPO":
		return Location{"Devonport", "Australia/Hobart"}
	case "DPS":
		return Location{"Denpasar-Bali Island", "Asia/Makassar"}
	case "DQA":
		return Location{"Daqing Shi", "Asia/Shanghai"}
	case "DQM":
		return Location{"Duqm", "Asia/Muscat"}
	case "DRB":
		return Location{"", "Australia/Perth"}
	case "DRG":
		return Location{"Deering", "America/Nome"}
	case "DRK":
		return Location{"Puntarenas", "America/Costa_Rica"}
	case "DRO":
		return Location{"Durango", "America/Denver"}
	case "DRS":
		return Location{"Dresden", "Europe/Berlin"}
	case "DRW":
		return Location{"Darwin", "Australia/Darwin"}
	case "DSE":
		return Location{"Dessie", "Africa/Addis_Ababa"}
	case "DSM":
		return Location{"Des Moines", "America/Chicago"}
	case "DSN":
		return Location{"Ordos", "Asia/Shanghai"}
	case "DSS":
		return Location{"Diass", "Africa/Dakar"}
	case "DTB":
		return Location{"Siborong-Borong", "Asia/Jakarta"}
	case "DTM":
		return Location{"Dortmund", "Europe/Berlin"}
	case "DTW":
		return Location{"Detroit", "America/Detroit"}
	case "DUB":
		return Location{"Dublin", "Europe/Dublin"}
	case "DUD":
		return Location{"Dunedin", "Pacific/Auckland"}
	case "DUE":
		return Location{"Chitato", "Africa/Luanda"}
	case "DUJ":
		return Location{"Dubois", "America/New_York"}
	case "DUR":
		return Location{"Durban", "Africa/Johannesburg"}
	case "DUS":
		return Location{"Dusseldorf", "Europe/Berlin"}
	case "DUT":
		return Location{"Unalaska", "America/Nome"}
	case "DVL":
		return Location{"Devils Lake", "America/Chicago"}
	case "DVO":
		return Location{"Davao City", "Asia/Manila"}
	case "DWC":
		return Location{"Jebel Ali", "Asia/Dubai"}
	case "DWD":
		return Location{"Dawadmi", "Asia/Riyadh"}
	case "DXB":
		return Location{"Dubai", "Asia/Dubai"}
	case "DYG":
		return Location{"Dayong", "Asia/Shanghai"}
	case "DYR":
		return Location{"Anadyr", "Asia/Anadyr"}
	case "DYU":
		return Location{"Dushanbe", "Asia/Dushanbe"}
	case "DZA":
		return Location{"Dzaoudzi", "Indian/Mayotte"}
	case "DZN":
		return Location{"Zhezkazgan", "Asia/Almaty"}
	case "EAA":
		return Location{"Eagle", "America/Anchorage"}
	case "EAE":
		return Location{"Sangafa", "Pacific/Efate"}
	case "EAM":
		return Location{"", "Asia/Riyadh"}
	case "EAR":
		return Location{"Kearney", "America/Chicago"}
	case "EAS":
		return Location{"Hondarribia", "Europe/Madrid"}
	case "EAT":
		return Location{"Wenatchee", "America/Los_Angeles"}
	case "EAU":
		return Location{"Eau Claire", "America/Chicago"}
	case "EBB":
		return Location{"Kampala", "Africa/Kampala"}
	case "EBH":
		return Location{"El Bayadh", "Africa/Algiers"}
	case "EBJ":
		return Location{"Esbjerg", "Europe/Copenhagen"}
	case "EBL":
		return Location{"Arbil", "Asia/Baghdad"}
	case "ECN":
		return Location{"Nicosia", "Asia/Famagusta"}
	case "ECP":
		return Location{"Panama City Beach", "America/Chicago"}
	case "EDI":
		return Location{"Edinburgh", "Europe/London"}
	case "EDL":
		return Location{"Eldoret", "Africa/Nairobi"}
	case "EDO":
		return Location{"Edremit", "Europe/Istanbul"}
	case "EDR":
		return Location{"", "Australia/Brisbane"}
	case "EEK":
		return Location{"Eek", "America/Nome"}
	case "EFL":
		return Location{"Kefallinia Island", "Europe/Athens"}
	case "EGC":
		return Location{"Bergerac/Roumaniere", "Europe/Paris"}
	case "EGE":
		return Location{"Eagle", "America/Denver"}
	case "EGM":
		return Location{"Sege", "Pacific/Guadalcanal"}
	case "EGS":
		return Location{"Egilsstadir", "Atlantic/Reykjavik"}
	case "EGX":
		return Location{"Egegik", "America/Anchorage"}
	case "EIN":
		return Location{"Eindhoven", "Europe/Amsterdam"}
	case "EIS":
		return Location{"Road Town", "America/Tortola"}
	case "EJA":
		return Location{"Barrancabermeja", "America/Bogota"}
	case "EKO":
		return Location{"Elko", "America/Los_Angeles"}
	case "ELC":
		return Location{"Elcho Island", "Australia/Darwin"}
	case "ELD":
		return Location{"El Dorado", "America/Chicago"}
	case "ELG":
		return Location{"", "Africa/Algiers"}
	case "ELH":
		return Location{"North Eleuthera", "America/Nassau"}
	case "ELI":
		return Location{"Elim", "America/Nome"}
	case "ELM":
		return Location{"Elmira/Corning", "America/New_York"}
	case "ELP":
		return Location{"El Paso", "America/Denver"}
	case "ELQ":
		return Location{"", "Asia/Riyadh"}
	case "ELS":
		return Location{"East London", "Africa/Johannesburg"}
	case "ELU":
		return Location{"Guemar", "Africa/Algiers"}
	case "EMA":
		return Location{"Nottingham", "Europe/London"}
	case "EMD":
		return Location{"Emerald", "Australia/Brisbane"}
	case "EMK":
		return Location{"Emmonak", "America/Nome"}
	case "EMN":
		return Location{"Nema", "Africa/Nouakchott"}
	case "ENA":
		return Location{"Kenai", "America/Anchorage"}
	case "ENE":
		return Location{"Ende-Flores Island", "Asia/Makassar"}
	case "ENH":
		return Location{"Enshi", "Asia/Shanghai"}
	case "ENI":
		return Location{"El Nido", "Asia/Manila"}
	case "ENU":
		return Location{"Enegu", "Africa/Lagos"}
	case "ENY":
		return Location{"Yan'an", "Asia/Shanghai"}
	case "EOH":
		return Location{"Medellin", "America/Bogota"}
	case "EPR":
		return Location{"", "Australia/Perth"}
	case "EQS":
		return Location{"Esquel", "America/Argentina/Catamarca"}
	case "ERC":
		return Location{"Erzincan", "Europe/Istanbul"}
	case "ERF":
		return Location{"Erfurt", "Europe/Berlin"}
	case "ERG":
		return Location{"Erbogachen", "Asia/Irkutsk"}
	case "ERH":
		return Location{"Errachidia", "Africa/Casablanca"}
	case "ERI":
		return Location{"Erie", "America/New_York"}
	case "ERL":
		return Location{"Erenhot", "Asia/Shanghai"}
	case "ERN":
		return Location{"Eirunepe", "America/Eirunepe"}
	case "ERS":
		return Location{"Windhoek", "Africa/Windhoek"}
	case "ERZ":
		return Location{"Erzurum", "Europe/Istanbul"}
	case "ESB":
		return Location{"Ankara", "Europe/Istanbul"}
	case "ESC":
		return Location{"Escanaba", "America/Detroit"}
	case "ESD":
		return Location{"Eastsound", "America/Los_Angeles"}
	case "ESU":
		return Location{"Essaouira", "Africa/Casablanca"}
	case "ETM":
		return Location{"Eilat", "Asia/Jerusalem"}
	case "ETR":
		return Location{"Santa Rosa", "America/Guayaquil"}
	case "ETZ":
		return Location{"Metz / Nancy", "Europe/Paris"}
	case "EUA":
		return Location{"Eua Island", "Pacific/Tongatapu"}
	case "EUG":
		return Location{"Eugene", "America/Los_Angeles"}
	case "EUN":
		return Location{"El Aaiun", "Africa/El_Aaiun"}
	case "EUQ":
		return Location{"San Jose", "Asia/Manila"}
	case "EUX":
		return Location{"Sint Eustatius", "America/Kralendijk"}
	case "EVE":
		return Location{"Evenes", "Europe/Oslo"}
	case "EVG":
		return Location{"", "Europe/Stockholm"}
	case "EVN":
		return Location{"Yerevan", "Asia/Yerevan"}
	case "EVV":
		return Location{"Evansville", "America/Chicago"}
	case "EWB":
		return Location{"New Bedford", "America/New_York"}
	case "EWN":
		return Location{"New Bern", "America/New_York"}
	case "EWR":
		return Location{"Newark", "America/New_York"}
	case "EXT":
		return Location{"Exeter", "Europe/London"}
	case "EYP":
		return Location{"El Yopal", "America/Bogota"}
	case "EYW":
		return Location{"Key West", "America/New_York"}
	case "EZE":
		return Location{"Ezeiza", "America/Argentina/Buenos_Aires"}
	case "EZS":
		return Location{"Elazig", "Europe/Istanbul"}
	case "EZV":
		return Location{"", "Asia/Yekaterinburg"}
	case "FAC":
		return Location{"", "Pacific/Tahiti"}
	case "FAE":
		return Location{"Vagar", "Atlantic/Faroe"}
	case "FAI":
		return Location{"Fairbanks", "America/Anchorage"}
	case "FAO":
		return Location{"Faro", "Europe/Lisbon"}
	case "FAR":
		return Location{"Fargo", "America/Chicago"}
	case "FAT":
		return Location{"Fresno", "America/Los_Angeles"}
	case "FAV":
		return Location{"", "Pacific/Tahiti"}
	case "FAY":
		return Location{"Fayetteville", "America/New_York"}
	case "FBE":
		return Location{"Francisco Beltrao", "America/Sao_Paulo"}
	case "FBM":
		return Location{"Lubumbashi", "Africa/Lubumbashi"}
	case "FCA":
		return Location{"Kalispell", "America/Denver"}
	case "FCO":
		return Location{"Rome", "Europe/Rome"}
	case "FDE":
		return Location{"Forde", "Europe/Oslo"}
	case "FDF":
		return Location{"Fort-de-France", "America/Martinique"}
	case "FDH":
		return Location{"Friedrichshafen", "Europe/Berlin"}
	case "FEC":
		return Location{"Feira De Santana", "America/Bahia"}
	case "FEG":
		return Location{"Fergana", "Asia/Tashkent"}
	case "FEN":
		return Location{"Fernando De Noronha", "America/Noronha"}
	case "FEZ":
		return Location{"Fes", "Africa/Casablanca"}
	case "FGU":
		return Location{"", "Pacific/Tahiti"}
	case "FHZ":
		return Location{"Fakahina", "Pacific/Tahiti"}
	case "FIH":
		return Location{"Kinshasa", "Africa/Kinshasa"}
	case "FIZ":
		return Location{"", "Australia/Perth"}
	case "FJR":
		return Location{"", "Asia/Dubai"}
	case "FKB":
		return Location{"Baden-Baden", "Europe/Berlin"}
	case "FKI":
		return Location{"Kisangani", "Africa/Lubumbashi"}
	case "FKQ":
		return Location{"Fakfak-Papua Island", "Asia/Jayapura"}
	case "FKS":
		return Location{"Sukagawa", "Asia/Tokyo"}
	case "FLA":
		return Location{"Florencia", "America/Bogota"}
	case "FLG":
		return Location{"Flagstaff", "America/Phoenix"}
	case "FLL":
		return Location{"Fort Lauderdale", "America/New_York"}
	case "FLN":
		return Location{"Florianopolis", "America/Sao_Paulo"}
	case "FLO":
		return Location{"Florence", "America/New_York"}
	case "FLR":
		return Location{"Firenze", "Europe/Rome"}
	case "FLS":
		return Location{"", "Australia/Hobart"}
	case "FLW":
		return Location{"Santa Cruz das Flores", "Atlantic/Azores"}
	case "FLZ":
		return Location{"Sibolga-Sumatra Island", "Asia/Jakarta"}
	case "FMA":
		return Location{"Formosa", "America/Argentina/Cordoba"}
	case "FMI":
		return Location{"", "Africa/Lubumbashi"}
	case "FMM":
		return Location{"Memmingen", "Europe/Berlin"}
	case "FMO":
		return Location{"Munster", "Europe/Berlin"}
	case "FNA":
		return Location{"Freetown", "Africa/Freetown"}
	case "FNC":
		return Location{"Funchal", "Atlantic/Madeira"}
	case "FNI":
		return Location{"Nimes/Garons", "Europe/Paris"}
	case "FNT":
		return Location{"Flint", "America/Detroit"}
	case "FOC":
		return Location{"Fuzhou", "Asia/Shanghai"}
	case "FOD":
		return Location{"Fort Dodge", "America/Chicago"}
	case "FOG":
		return Location{"Foggia", "Europe/Rome"}
	case "FON":
		return Location{"La Fortuna/San Carlos", "America/Costa_Rica"}
	case "FOR":
		return Location{"Fortaleza", "America/Fortaleza"}
	case "FPO":
		return Location{"Freeport", "America/Nassau"}
	case "FRA":
		return Location{"Frankfurt-am-Main", "Europe/Berlin"}
	case "FRD":
		return Location{"Friday Harbor", "America/Los_Angeles"}
	case "FRE":
		return Location{"Fera Island", "Pacific/Guadalcanal"}
	case "FRL":
		return Location{"Forli", "Europe/Rome"}
	case "FRO":
		return Location{"Floro", "Europe/Oslo"}
	case "FRS":
		return Location{"San Benito", "America/Guatemala"}
	case "FRU":
		return Location{"Bishkek", "Asia/Bishkek"}
	case "FRW":
		return Location{"Francistown", "Africa/Gaborone"}
	case "FSC":
		return Location{"Figari Sud-Corse", "Europe/Paris"}
	case "FSD":
		return Location{"Sioux Falls", "America/Chicago"}
	case "FSM":
		return Location{"Fort Smith", "America/Chicago"}
	case "FSZ":
		return Location{"", "Asia/Tokyo"}
	case "FTA":
		return Location{"Futuna Island", "Pacific/Efate"}
	case "FTE":
		return Location{"El Calafate", "America/Argentina/Rio_Gallegos"}
	case "FTU":
		return Location{"Tolanaro", "Indian/Antananarivo"}
	case "FUE":
		return Location{"Fuerteventura Island", "Atlantic/Canary"}
	case "FUG":
		return Location{"Fuyang", "Asia/Shanghai"}
	case "FUJ":
		return Location{"Goto", "Asia/Tokyo"}
	case "FUK":
		return Location{"Fukuoka", "Asia/Tokyo"}
	case "FUN":
		return Location{"Funafuti", "Pacific/Funafuti"}
	case "FUO":
		return Location{"Foshan", "Asia/Shanghai"}
	case "FVM":
		return Location{"Fuvahmulah Island", "Indian/Maldives"}
	case "FWA":
		return Location{"Fort Wayne", "America/Indiana/Indianapolis"}
	case "FYU":
		return Location{"Fort Yukon", "America/Anchorage"}
	case "GAE":
		return Location{"Gabes", "Africa/Tunis"}
	case "GAJ":
		return Location{"Yamagata", "Asia/Tokyo"}
	case "GAL":
		return Location{"Galena", "America/Anchorage"}
	case "GAM":
		return Location{"Gambell", "America/Nome"}
	case "GAN":
		return Location{"Gan", "Indian/Maldives"}
	case "GAU":
		return Location{"Guwahati", "Asia/Kolkata"}
	case "GAY":
		return Location{"", "Asia/Kolkata"}
	case "GBE":
		return Location{"Gaborone", "Africa/Gaborone"}
	case "GBI":
		return Location{"Grand Bahama", "America/Nassau"}
	case "GBT":
		return Location{"Gorgan", "Asia/Tehran"}
	case "GCC":
		return Location{"Gillette", "America/Denver"}
	case "GCI":
		return Location{"Saint Peter Port", "Europe/Guernsey"}
	case "GCK":
		return Location{"Garden City", "America/Chicago"}
	case "GCM":
		return Location{"Georgetown", "America/Cayman"}
	case "GCN":
		return Location{"Grand Canyon", "America/Phoenix"}
	case "GCW":
		return Location{"Peach Springs", "America/Phoenix"}
	case "GDE":
		return Location{"Gode", "Africa/Addis_Ababa"}
	case "GDL":
		return Location{"Guadalajara", "America/Mexico_City"}
	case "GDN":
		return Location{"Gdansk", "Europe/Warsaw"}
	case "GDQ":
		return Location{"Gondar", "Africa/Addis_Ababa"}
	case "GDT":
		return Location{"Cockburn Town", "America/Grand_Turk"}
	case "GDV":
		return Location{"Glendive", "America/Denver"}
	case "GDX":
		return Location{"Magadan", "Asia/Magadan"}
	case "GEA":
		return Location{"Noumea", "Pacific/Noumea"}
	case "GEG":
		return Location{"Spokane", "America/Los_Angeles"}
	case "GEL":
		return Location{"Santo Angelo", "America/Sao_Paulo"}
	case "GEO":
		return Location{"Georgetown", "America/Guyana"}
	case "GES":
		return Location{"General Santos City", "Asia/Manila"}
	case "GET":
		return Location{"", "Australia/Perth"}
	case "GEV":
		return Location{"Gallivare", "Europe/Stockholm"}
	case "GFF":
		return Location{"Griffith", "Australia/Sydney"}
	case "GFK":
		return Location{"Grand Forks", "America/Chicago"}
	case "GGF":
		return Location{"Almeirim", "America/Santarem"}
	case "GGG":
		return Location{"Longview", "America/Chicago"}
	case "GGM":
		return Location{"Kakamega", "Africa/Nairobi"}
	case "GGT":
		return Location{"George Town", "America/Nassau"}
	case "GGW":
		return Location{"Glasgow", "America/Denver"}
	case "GHA":
		return Location{"Ghardaia", "Africa/Algiers"}
	case "GHB":
		return Location{"Governor's Harbour", "America/Nassau"}
	case "GHC":
		return Location{"", "America/Nassau"}
	case "GHT":
		return Location{"Ghat", "Africa/Tripoli"}
	case "GIB":
		return Location{"Gibraltar", "Europe/Gibraltar"}
	case "GIC":
		return Location{"", "Australia/Brisbane"}
	case "GIG":
		return Location{"Rio De Janeiro", "America/Sao_Paulo"}
	case "GIL":
		return Location{"Gilgit", "Asia/Karachi"}
	case "GIS":
		return Location{"Gisborne", "Pacific/Auckland"}
	case "GIU":
		return Location{"Sigiriya", "Asia/Colombo"}
	case "GIZ":
		return Location{"Jizan", "Asia/Riyadh"}
	case "GJA":
		return Location{"Guanaja", "America/Tegucigalpa"}
	case "GJL":
		return Location{"Jijel", "Africa/Algiers"}
	case "GJT":
		return Location{"Grand Junction", "America/Denver"}
	case "GKA":
		return Location{"Goronka", "Pacific/Port_Moresby"}
	case "GLA":
		return Location{"Glasgow", "Europe/London"}
	case "GLF":
		return Location{"Golfito", "America/Costa_Rica"}
	case "GLH":
		return Location{"Greenville", "America/Chicago"}
	case "GLT":
		return Location{"Gladstone", "Australia/Brisbane"}
	case "GLV":
		return Location{"Golovin", "America/Nome"}
	case "GMA":
		return Location{"Gemena", "Africa/Kinshasa"}
	case "GMB":
		return Location{"Gambela", "Africa/Addis_Ababa"}
	case "GMO":
		return Location{"Gombe", "Africa/Lagos"}
	case "GMP":
		return Location{"Seoul", "Asia/Seoul"}
	case "GMR":
		return Location{"", "Pacific/Gambier"}
	case "GMZ":
		return Location{"Alajero", "Atlantic/Canary"}
	case "GNB":
		return Location{"Grenoble/Saint-Geoirs", "Europe/Paris"}
	case "GND":
		return Location{"Saint George's", "America/Grenada"}
	case "GNJ":
		return Location{"Ganja", "Asia/Baku"}
	case "GNM":
		return Location{"Guanambi", "America/Bahia"}
	case "GNS":
		return Location{"Gunung Sitoli-Nias Island", "Asia/Jakarta"}
	case "GNV":
		return Location{"Gainesville", "America/New_York"}
	case "GNY":
		return Location{"Sanliurfa", "Europe/Istanbul"}
	case "GOA":
		return Location{"Genova", "Europe/Rome"}
	case "GOB":
		return Location{"Goba", "Africa/Addis_Ababa"}
	case "GOH":
		return Location{"Nuuk", "America/Nuuk"}
	case "GOI":
		return Location{"Dabolim", "Asia/Kolkata"}
	case "GOJ":
		return Location{"Nizhny Novgorod", "Europe/Moscow"}
	case "GOM":
		return Location{"Goma", "Africa/Kigali"}
	case "GOP":
		return Location{"Gorakhpur", "Asia/Kolkata"}
	case "GOQ":
		return Location{"Golmud", "Asia/Shanghai"}
	case "GOT":
		return Location{"Gothenburg", "Europe/Stockholm"}
	case "GOU":
		return Location{"Garoua", "Africa/Douala"}
	case "GOV":
		return Location{"Nhulunbuy", "Australia/Darwin"}
	case "GOY":
		return Location{"Amparai", "Asia/Colombo"}
	case "GPA":
		return Location{"Patras", "Europe/Athens"}
	case "GPB":
		return Location{"Guarapuava", "America/Sao_Paulo"}
	case "GPI":
		return Location{"Guapi", "America/Bogota"}
	case "GPS":
		return Location{"Baltra", "Pacific/Galapagos"}
	case "GPT":
		return Location{"Gulfport", "America/Chicago"}
	case "GRB":
		return Location{"Green Bay", "America/Chicago"}
	case "GRI":
		return Location{"Grand Island", "America/Chicago"}
	case "GRJ":
		return Location{"George", "Africa/Johannesburg"}
	case "GRK":
		return Location{"Fort Hood/Killeen", "America/Chicago"}
	case "GRO":
		return Location{"Girona", "Europe/Madrid"}
	case "GRQ":
		return Location{"Groningen", "Europe/Amsterdam"}
	case "GRR":
		return Location{"Grand Rapids", "America/Detroit"}
	case "GRU":
		return Location{"Sao Paulo", "America/Sao_Paulo"}
	case "GRV":
		return Location{"Grozny", "Europe/Moscow"}
	case "GRW":
		return Location{"Santa Cruz da Graciosa", "Atlantic/Azores"}
	case "GRX":
		return Location{"Granada", "Europe/Madrid"}
	case "GRZ":
		return Location{"Graz", "Europe/Vienna"}
	case "GSM":
		return Location{"", "Asia/Tehran"}
	case "GSO":
		return Location{"Greensboro", "America/New_York"}
	case "GSP":
		return Location{"Greenville", "America/New_York"}
	case "GST":
		return Location{"Gustavus", "America/Juneau"}
	case "GSV":
		return Location{"Saratov", "Europe/Saratov"}
	case "GTE":
		return Location{"Groote Eylandt", "Australia/Darwin"}
	case "GTF":
		return Location{"Great Falls", "America/Denver"}
	case "GTO":
		return Location{"Gorontalo-Celebes Island", "Asia/Makassar"}
	case "GTP":
		return Location{"Grants Pass", "America/Los_Angeles"}
	case "GTR":
		return Location{"Columbus/W Point/Starkville", "America/Chicago"}
	case "GTS":
		return Location{"", "Australia/Adelaide"}
	case "GUA":
		return Location{"Guatemala City", "America/Guatemala"}
	case "GUC":
		return Location{"Gunnison", "America/Denver"}
	case "GUM":
		return Location{"Hagatna", "Pacific/Guam"}
	case "GUP":
		return Location{"Gallup", "America/Denver"}
	case "GUR":
		return Location{"Gurney", "Pacific/Port_Moresby"}
	case "GUW":
		return Location{"Atyrau", "Asia/Atyrau"}
	case "GVA":
		return Location{"Geneva", "Europe/Paris"}
	case "GVR":
		return Location{"Governador Valadares", "America/Sao_Paulo"}
	case "GWD":
		return Location{"Gwadar", "Asia/Karachi"}
	case "GWL":
		return Location{"Gwalior", "Asia/Kolkata"}
	case "GWT":
		return Location{"Westerland", "Europe/Berlin"}
	case "GXF":
		return Location{"Sayun", "Asia/Aden"}
	case "GYD":
		return Location{"Baku", "Asia/Baku"}
	case "GYE":
		return Location{"Guayaquil", "America/Guayaquil"}
	case "GYN":
		return Location{"Goiania", "America/Sao_Paulo"}
	case "GYS":
		return Location{"Guangyuan", "Asia/Shanghai"}
	case "GYU":
		return Location{"Guyuan", "Asia/Shanghai"}
	case "GZO":
		return Location{"Gizo", "Pacific/Guadalcanal"}
	case "GZP":
		return Location{"Gazipasa", "Europe/Istanbul"}
	case "GZT":
		return Location{"Gaziantep", "Europe/Istanbul"}
	case "HAA":
		return Location{"Hasvik", "Europe/Oslo"}
	case "HAC":
		return Location{"Hachijojima", "Asia/Tokyo"}
	case "HAD":
		return Location{"Halmstad", "Europe/Stockholm"}
	case "HAH":
		return Location{"Moroni", "Indian/Comoro"}
	case "HAJ":
		return Location{"Hannover", "Europe/Berlin"}
	case "HAK":
		return Location{"Haikou", "Asia/Shanghai"}
	case "HAM":
		return Location{"Hamburg", "Europe/Berlin"}
	case "HAN":
		return Location{"Hanoi", "Asia/Bangkok"}
	case "HAQ":
		return Location{"Haa Dhaalu Atoll", "Indian/Maldives"}
	case "HAS":
		return Location{"", "Asia/Riyadh"}
	case "HAU":
		return Location{"Karmoy", "Europe/Oslo"}
	case "HAV":
		return Location{"Havana", "America/Havana"}
	case "HAY":
		return Location{"Aguachica", "America/Bogota"}
	case "HBA":
		return Location{"Hobart", "Australia/Hobart"}
	case "HBE":
		return Location{"Alexandria", "Africa/Cairo"}
	case "HBX":
		return Location{"Hubli", "Asia/Kolkata"}
	case "HCQ":
		return Location{"", "Australia/Perth"}
	case "HCR":
		return Location{"Holy Cross", "America/Anchorage"}
	case "HDF":
		return Location{"Heringsdorf", "Europe/Berlin"}
	case "HDG":
		return Location{"Handan", "Asia/Shanghai"}
	case "HDK":
		return Location{"Kulhudhuffushi", "Indian/Maldives"}
	case "HDN":
		return Location{"Hayden", "America/Denver"}
	case "HDS":
		return Location{"Hoedspruit", "Africa/Johannesburg"}
	case "HDY":
		return Location{"Hat Yai", "Asia/Bangkok"}
	case "HEA":
		return Location{"", "Asia/Kabul"}
	case "HEH":
		return Location{"Heho", "Asia/Yangon"}
	case "HEK":
		return Location{"Heihe", "Asia/Shanghai"}
	case "HEL":
		return Location{"Helsinki", "Europe/Helsinki"}
	case "HER":
		return Location{"Heraklion", "Europe/Athens"}
	case "HET":
		return Location{"Hohhot", "Asia/Shanghai"}
	case "HFE":
		return Location{"Hefei", "Asia/Shanghai"}
	case "HFS":
		return Location{"", "Europe/Stockholm"}
	case "HFT":
		return Location{"Hammerfest", "Europe/Oslo"}
	case "HGA":
		return Location{"Hargeisa", "Africa/Mogadishu"}
	case "HGD":
		return Location{"", "Australia/Brisbane"}
	case "HGH":
		return Location{"Hangzhou", "Asia/Shanghai"}
	case "HGN":
		return Location{"", "Asia/Bangkok"}
	case "HGO":
		return Location{"", "Africa/Abidjan"}
	case "HGR":
		return Location{"Hagerstown", "America/New_York"}
	case "HGU":
		return Location{"Mount Hagen", "Pacific/Port_Moresby"}
	case "HHH":
		return Location{"Hilton Head Island", "America/New_York"}
	case "HHN":
		return Location{"Hahn", "Europe/Berlin"}
	case "HHQ":
		return Location{"Hua Hin", "Asia/Bangkok"}
	case "HHR":
		return Location{"Hawthorne", "America/Los_Angeles"}
	case "HHZ":
		return Location{"Hikueru Atoll", "Pacific/Tahiti"}
	case "HIA":
		return Location{"Huai'an", "Asia/Shanghai"}
	case "HIB":
		return Location{"Hibbing", "America/Chicago"}
	case "HID":
		return Location{"Horn Island", "Australia/Brisbane"}
	case "HIJ":
		return Location{"Hiroshima", "Asia/Tokyo"}
	case "HIN":
		return Location{"Sacheon", "Asia/Seoul"}
	case "HIR":
		return Location{"Honiara", "Pacific/Guadalcanal"}
	case "HJJ":
		return Location{"Huaihua", "Asia/Shanghai"}
	case "HJR":
		return Location{"Khajuraho", "Asia/Kolkata"}
	case "HKD":
		return Location{"Hakodate", "Asia/Tokyo"}
	case "HKG":
		return Location{"Hong Kong", "Asia/Hong_Kong"}
	case "HKK":
		return Location{"", "Pacific/Auckland"}
	case "HKN":
		return Location{"Hoskins", "Pacific/Port_Moresby"}
	case "HKT":
		return Location{"Phuket", "Asia/Bangkok"}
	case "HLA":
		return Location{"Johannesburg", "Africa/Johannesburg"}
	case "HLD":
		return Location{"Hailar", "Asia/Shanghai"}
	case "HLH":
		return Location{"Ulanhot", "Asia/Shanghai"}
	case "HLN":
		return Location{"Helena", "America/Denver"}
	case "HLP":
		return Location{"Jakarta", "Asia/Jakarta"}
	case "HLZ":
		return Location{"Hamilton", "Pacific/Auckland"}
	case "HMA":
		return Location{"Khanty-Mansiysk", "Asia/Yekaterinburg"}
	case "HMB":
		return Location{"Sohag", "Africa/Cairo"}
	case "HME":
		return Location{"Hassi Messaoud", "Africa/Algiers"}
	case "HMI":
		return Location{"Hami", "Asia/Shanghai"}
	case "HMO":
		return Location{"Hermosillo", "America/Hermosillo"}
	case "HMV":
		return Location{"", "Europe/Stockholm"}
	case "HNA":
		return Location{"", "Asia/Tokyo"}
	case "HND":
		return Location{"Tokyo", "Asia/Tokyo"}
	case "HNH":
		return Location{"Hoonah", "America/Juneau"}
	case "HNL":
		return Location{"Honolulu", "Pacific/Honolulu"}
	case "HNM":
		return Location{"Hana", "Pacific/Honolulu"}
	case "HNS":
		return Location{"Haines", "America/Juneau"}
	case "HNY":
		return Location{"Hengyang", "Asia/Shanghai"}
	case "HOB":
		return Location{"Hobbs", "America/Denver"}
	case "HOF":
		return Location{"", "Asia/Riyadh"}
	case "HOG":
		return Location{"Holguin", "America/Havana"}
	case "HOI":
		return Location{"", "Pacific/Tahiti"}
	case "HOM":
		return Location{"Homer", "America/Anchorage"}
	case "HOR":
		return Location{"Horta", "Atlantic/Azores"}
	case "HOT":
		return Location{"Hot Springs", "America/Chicago"}
	case "HOU":
		return Location{"Houston", "America/Chicago"}
	case "HOV":
		return Location{"Orsta", "Europe/Oslo"}
	case "HOX":
		return Location{"Hommalinn", "Asia/Yangon"}
	case "HPA":
		return Location{"Lifuka", "Pacific/Tongatapu"}
	case "HPB":
		return Location{"Hooper Bay", "America/Nome"}
	case "HPH":
		return Location{"Haiphong", "Asia/Bangkok"}
	case "HPN":
		return Location{"White Plains", "America/New_York"}
	case "HRB":
		return Location{"Harbin", "Asia/Shanghai"}
	case "HRE":
		return Location{"Harare", "Africa/Harare"}
	case "HRG":
		return Location{"Hurghada", "Africa/Cairo"}
	case "HRL":
		return Location{"Harlingen", "America/Chicago"}
	case "HRO":
		return Location{"Harrison", "America/Chicago"}
	case "HSA":
		return Location{"Turkistan", "Asia/Almaty"}
	case "HSG":
		return Location{"Saga", "Asia/Tokyo"}
	case "HSL":
		return Location{"Huslia", "America/Anchorage"}
	case "HSN":
		return Location{"Zhoushan", "Asia/Shanghai"}
	case "HSV":
		return Location{"Huntsville", "America/Chicago"}
	case "HTA":
		return Location{"Chita", "Asia/Chita"}
	case "HTG":
		return Location{"Khatanga", "Asia/Krasnoyarsk"}
	case "HTI":
		return Location{"Hamilton Island", "Australia/Lindeman"}
	case "HTN":
		return Location{"Hotan", "Asia/Shanghai"}
	case "HTS":
		return Location{"Huntington", "America/New_York"}
	case "HTY":
		return Location{"Hatay", "Europe/Istanbul"}
	case "HUH":
		return Location{"Fare", "Pacific/Tahiti"}
	case "HUI":
		return Location{"Hue", "Asia/Ho_Chi_Minh"}
	case "HUN":
		return Location{"Hualien City", "Asia/Taipei"}
	case "HUS":
		return Location{"Hughes", "America/Anchorage"}
	case "HUU":
		return Location{"Huanuco", "America/Lima"}
	case "HUX":
		return Location{"Huatulco", "America/Mexico_City"}
	case "HUY":
		return Location{"Grimsby", "Europe/London"}
	case "HUZ":
		return Location{"Huizhou", "Asia/Shanghai"}
	case "HVB":
		return Location{"Hervey Bay", "Australia/Brisbane"}
	case "HVD":
		return Location{"Khovd", "Asia/Hovd"}
	case "HVG":
		return Location{"Honningsvag", "Europe/Oslo"}
	case "HVN":
		return Location{"New Haven", "America/New_York"}
	case "HVR":
		return Location{"Havre", "America/Denver"}
	case "HWN":
		return Location{"Hwange", "Africa/Harare"}
	case "HYA":
		return Location{"Hyannis", "America/New_York"}
	case "HYD":
		return Location{"Hyderabad", "Asia/Kolkata"}
	case "HYN":
		return Location{"Huangyan", "Asia/Shanghai"}
	case "HYS":
		return Location{"Hays", "America/Chicago"}
	case "HZG":
		return Location{"Hanzhong", "Asia/Shanghai"}
	case "IAA":
		return Location{"Igarka", "Asia/Krasnoyarsk"}
	case "IAD":
		return Location{"Dulles", "America/New_York"}
	case "IAG":
		return Location{"Niagara Falls", "America/New_York"}
	case "IAH":
		return Location{"Houston", "America/Chicago"}
	case "IAM":
		return Location{"Amenas", "Africa/Algiers"}
	case "IAN":
		return Location{"Kiana", "America/Anchorage"}
	case "IAO":
		return Location{"Del Carmen", "Asia/Manila"}
	case "IAR":
		return Location{"", "Europe/Moscow"}
	case "IAS":
		return Location{"Iasi", "Europe/Bucharest"}
	case "IBA":
		return Location{"Ibadan", "Africa/Lagos"}
	case "IBE":
		return Location{"Ibague", "America/Bogota"}
	case "IBR":
		return Location{"Omitama", "Asia/Tokyo"}
	case "IBZ":
		return Location{"Ibiza", "Europe/Madrid"}
	case "ICI":
		return Location{"Cicia", "Pacific/Fiji"}
	case "ICN":
		return Location{"Seoul", "Asia/Seoul"}
	case "ICT":
		return Location{"Wichita", "America/Chicago"}
	case "IDA":
		return Location{"Idaho Falls", "America/Boise"}
	case "IDR":
		return Location{"Indore", "Asia/Kolkata"}
	case "IEG":
		return Location{"Babimost", "Europe/Warsaw"}
	case "IEV":
		return Location{"Kiev", "Europe/Kiev"}
	case "IFJ":
		return Location{"Isafjordur", "Atlantic/Reykjavik"}
	case "IFN":
		return Location{"Isfahan", "Asia/Tehran"}
	case "IFU":
		return Location{"Ifuru", "Indian/Maldives"}
	case "IGA":
		return Location{"Matthew Town", "America/Nassau"}
	case "IGD":
		return Location{"Igdir", "Europe/Istanbul"}
	case "IGG":
		return Location{"Igiugig", "America/Anchorage"}
	case "IGR":
		return Location{"Puerto Iguazu", "America/Argentina/Cordoba"}
	case "IGT":
		return Location{"Magas", "Europe/Moscow"}
	case "IGU":
		return Location{"Foz Do Iguacu", "America/Argentina/Cordoba"}
	case "IIL":
		return Location{"Ilam", "Asia/Tehran"}
	case "IJK":
		return Location{"Izhevsk", "Europe/Samara"}
	case "IKA":
		return Location{"Tehran", "Asia/Tehran"}
	case "IKI":
		return Location{"Iki", "Asia/Tokyo"}
	case "IKO":
		return Location{"Nikolski", "America/Nome"}
	case "IKS":
		return Location{"Tiksi", "Asia/Yakutsk"}
	case "IKT":
		return Location{"Irkutsk", "Asia/Irkutsk"}
	case "IKU":
		return Location{"Tamchy", "Asia/Bishkek"}
	case "ILD":
		return Location{"Lleida", "Europe/Madrid"}
	case "ILG":
		return Location{"Wilmington", "America/New_York"}
	case "ILI":
		return Location{"Iliamna", "America/Anchorage"}
	case "ILM":
		return Location{"Wilmington", "America/New_York"}
	case "ILO":
		return Location{"Iloilo City", "Asia/Manila"}
	case "ILP":
		return Location{"Ile des Pins", "Pacific/Noumea"}
	case "ILQ":
		return Location{"Ilo", "America/Lima"}
	case "ILR":
		return Location{"Ilorin", "Africa/Lagos"}
	case "ILY":
		return Location{"Port Ellen", "Europe/London"}
	case "IMF":
		return Location{"Imphal", "Asia/Kolkata"}
	case "IMP":
		return Location{"Imperatriz", "America/Fortaleza"}
	case "IMT":
		return Location{"Iron Mountain Kingsford", "America/Chicago"}
	case "INC":
		return Location{"Yinchuan", "Asia/Shanghai"}
	case "IND":
		return Location{"Indianapolis", "America/Indiana/Indianapolis"}
	case "INF":
		return Location{"In Guezzam", "Africa/Algiers"}
	case "INH":
		return Location{"Inhambabe", "Africa/Maputo"}
	case "INI":
		return Location{"Nis", "Europe/Belgrade"}
	case "INL":
		return Location{"International Falls", "America/Chicago"}
	case "INN":
		return Location{"Innsbruck", "Europe/Vienna"}
	case "INU":
		return Location{"Yaren District", "Pacific/Nauru"}
	case "INV":
		return Location{"Inverness", "Europe/London"}
	case "INZ":
		return Location{"In Salah", "Africa/Algiers"}
	case "IOA":
		return Location{"Ioannina", "Europe/Athens"}
	case "IOM":
		return Location{"Castletown", "Europe/Isle_of_Man"}
	case "IOS":
		return Location{"Ilheus", "America/Bahia"}
	case "IPA":
		return Location{"Ipota", "Pacific/Efate"}
	case "IPC":
		return Location{"Isla De Pascua", "Pacific/Easter"}
	case "IPH":
		return Location{"Ipoh", "Asia/Kuala_Lumpur"}
	case "IPI":
		return Location{"Ipiales", "America/Bogota"}
	case "IPL":
		return Location{"Imperial", "America/Los_Angeles"}
	case "IPN":
		return Location{"Ipatinga", "America/Sao_Paulo"}
	case "IQM":
		return Location{"Qiemo", "Asia/Shanghai"}
	case "IQN":
		return Location{"Qingyang", "Asia/Shanghai"}
	case "IQQ":
		return Location{"Iquique", "America/Santiago"}
	case "IQT":
		return Location{"Iquitos", "America/Lima"}
	case "IRA":
		return Location{"Kirakira", "Pacific/Guadalcanal"}
	case "IRC":
		return Location{"Circle", "America/Anchorage"}
	case "IRG":
		return Location{"", "Australia/Brisbane"}
	case "IRI":
		return Location{"Nduli", "Africa/Dar_es_Salaam"}
	case "IRJ":
		return Location{"La Rioja", "America/Argentina/La_Rioja"}
	case "IRK":
		return Location{"Kirksville", "America/Chicago"}
	case "IRM":
		return Location{"", "Asia/Yekaterinburg"}
	case "IRP":
		return Location{"", "Africa/Lubumbashi"}
	case "IRZ":
		return Location{"Santa Isabel Do Rio Negro", "America/Manaus"}
	case "ISA":
		return Location{"Mount Isa", "Australia/Brisbane"}
	case "ISB":
		return Location{"Islamabad", "Asia/Karachi"}
	case "ISE":
		return Location{"Isparta", "Europe/Istanbul"}
	case "ISG":
		return Location{"Ishigaki", "Asia/Tokyo"}
	case "ISK":
		return Location{"Nashik", "Asia/Kolkata"}
	case "ISP":
		return Location{"Islip", "America/New_York"}
	case "IST":
		return Location{"Arnavutkoy", "Europe/Istanbul"}
	case "ISU":
		return Location{"Sulaymaniyah", "Asia/Baghdad"}
	case "ITB":
		return Location{"Itaituba", "America/Santarem"}
	case "ITH":
		return Location{"Ithaca", "America/New_York"}
	case "ITM":
		return Location{"Osaka", "Asia/Tokyo"}
	case "ITO":
		return Location{"Hilo", "Pacific/Honolulu"}
	case "IUE":
		return Location{"Alofi", "Pacific/Niue"}
	case "IVC":
		return Location{"Invercargill", "Pacific/Auckland"}
	case "IVL":
		return Location{"Ivalo", "Europe/Helsinki"}
	case "IVR":
		return Location{"", "Australia/Sydney"}
	case "IWA":
		return Location{"Ivanovo", "Europe/Moscow"}
	case "IWD":
		return Location{"Ironwood", "America/Menominee"}
	case "IWJ":
		return Location{"Masuda", "Asia/Tokyo"}
	case "IWK":
		return Location{"Iwakuni", "Asia/Tokyo"}
	case "IXA":
		return Location{"Agartala", "Asia/Dhaka"}
	case "IXB":
		return Location{"Siliguri", "Asia/Kolkata"}
	case "IXC":
		return Location{"Chandigarh", "Asia/Kolkata"}
	case "IXD":
		return Location{"Allahabad", "Asia/Kolkata"}
	case "IXE":
		return Location{"Mangalore", "Asia/Kolkata"}
	case "IXG":
		return Location{"", "Asia/Kolkata"}
	case "IXI":
		return Location{"Lilabari", "Asia/Kolkata"}
	case "IXJ":
		return Location{"Jammu", "Asia/Kolkata"}
	case "IXK":
		return Location{"", "Asia/Kolkata"}
	case "IXL":
		return Location{"Leh", "Asia/Kolkata"}
	case "IXM":
		return Location{"Madurai", "Asia/Kolkata"}
	case "IXR":
		return Location{"Ranchi", "Asia/Kolkata"}
	case "IXS":
		return Location{"Silchar", "Asia/Kolkata"}
	case "IXT":
		return Location{"Pasighat", "Asia/Kolkata"}
	case "IXU":
		return Location{"Aurangabad", "Asia/Kolkata"}
	case "IXW":
		return Location{"", "Asia/Kolkata"}
	case "IXY":
		return Location{"Kandla", "Asia/Kolkata"}
	case "IXZ":
		return Location{"Port Blair", "Asia/Kolkata"}
	case "IZA":
		return Location{"Juiz De Fora", "America/Sao_Paulo"}
	case "IZO":
		return Location{"Izumo", "Asia/Tokyo"}
	case "JAC":
		return Location{"Jackson", "America/Denver"}
	case "JAE":
		return Location{"Jaen", "America/Lima"}
	case "JAF":
		return Location{"Jaffna", "Asia/Colombo"}
	case "JAI":
		return Location{"Jaipur", "Asia/Kolkata"}
	case "JAN":
		return Location{"Jackson", "America/Chicago"}
	case "JAU":
		return Location{"Jauja", "America/Lima"}
	case "JAV":
		return Location{"Ilulissat", "America/Nuuk"}
	case "JAX":
		return Location{"Jacksonville", "America/New_York"}
	case "JBQ":
		return Location{"La Isabela", "America/Santo_Domingo"}
	case "JBR":
		return Location{"Jonesboro", "America/Chicago"}
	case "JCK":
		return Location{"", "Australia/Brisbane"}
	case "JDH":
		return Location{"Jodhpur", "Asia/Kolkata"}
	case "JDO":
		return Location{"Juazeiro Do Norte", "America/Fortaleza"}
	case "JDZ":
		return Location{"Jingdezhen", "Asia/Shanghai"}
	case "JED":
		return Location{"Jeddah", "Asia/Riyadh"}
	case "JEE":
		return Location{"Jeremie", "America/Port-au-Prince"}
	case "JEG":
		return Location{"Aasiaat", "America/Nuuk"}
	case "JEK":
		return Location{"Lower Zambezi National Park", "Africa/Harare"}
	case "JER":
		return Location{"Saint Helier", "Europe/Jersey"}
	case "JFK":
		return Location{"New York", "America/New_York"}
	case "JFR":
		return Location{"Paamiut", "America/Nuuk"}
	case "JGA":
		return Location{"Jamnagar", "Asia/Kolkata"}
	case "JGN":
		return Location{"Jiayuguan", "Asia/Shanghai"}
	case "JGS":
		return Location{"Ji'an", "Asia/Shanghai"}
	case "JHB":
		return Location{"Senai", "Asia/Kuala_Lumpur"}
	case "JHG":
		return Location{"Jinghong", "Asia/Shanghai"}
	case "JHM":
		return Location{"Lahaina", "Pacific/Honolulu"}
	case "JHS":
		return Location{"Sisimiut", "America/Nuuk"}
	case "JIA":
		return Location{"Juina", "America/Cuiaba"}
	case "JIB":
		return Location{"Djibouti City", "Africa/Djibouti"}
	case "JIC":
		return Location{"Jinchang", "Asia/Shanghai"}
	case "JIJ":
		return Location{"Jijiga", "Africa/Addis_Ababa"}
	case "JIK":
		return Location{"Ikaria Island", "Europe/Athens"}
	case "JIM":
		return Location{"Jimma", "Africa/Addis_Ababa"}
	case "JIN":
		return Location{"Jinja", "Africa/Kampala"}
	case "JIQ":
		return Location{"Chongqing", "Asia/Shanghai"}
	case "JIU":
		return Location{"Jiujiang", "Asia/Shanghai"}
	case "JJD":
		return Location{"Jijoca de Jericoacoara", "America/Fortaleza"}
	case "JJN":
		return Location{"Quanzhou", "Asia/Shanghai"}
	case "JKH":
		return Location{"Chios Island", "Europe/Athens"}
	case "JKL":
		return Location{"Kalymnos Island", "Europe/Athens"}
	case "JKR":
		return Location{"Janakpur", "Asia/Kathmandu"}
	case "JLN":
		return Location{"Joplin", "America/Chicago"}
	case "JLR":
		return Location{"", "Asia/Kolkata"}
	case "JMK":
		return Location{"Mykonos Island", "Europe/Athens"}
	case "JMS":
		return Location{"Jamestown", "America/Chicago"}
	case "JMU":
		return Location{"Jiamusi", "Asia/Shanghai"}
	case "JNB":
		return Location{"Johannesburg", "Africa/Johannesburg"}
	case "JNG":
		return Location{"Jining", "Asia/Shanghai"}
	case "JNU":
		return Location{"Juneau", "America/Juneau"}
	case "JNX":
		return Location{"Naxos Island", "Europe/Athens"}
	case "JNZ":
		return Location{"Jinzhou", "Asia/Shanghai"}
	case "JOE":
		return Location{"Joensuu / Liperi", "Europe/Helsinki"}
	case "JOG":
		return Location{"Yogyakarta-Java Island", "Asia/Jakarta"}
	case "JOI":
		return Location{"Joinville", "America/Sao_Paulo"}
	case "JOS":
		return Location{"Jos", "Africa/Lagos"}
	case "JPA":
		return Location{"Joao Pessoa", "America/Fortaleza"}
	case "JPR":
		return Location{"Ji-Parana", "America/Porto_Velho"}
	case "JQA":
		return Location{"Uummannaq", "America/Nuuk"}
	case "JRG":
		return Location{"Jharsuguda", "Asia/Kolkata"}
	case "JRH":
		return Location{"Jorhat", "Asia/Kolkata"}
	case "JRO":
		return Location{"Arusha", "Africa/Dar_es_Salaam"}
	case "JSA":
		return Location{"", "Asia/Kolkata"}
	case "JSH":
		return Location{"Crete Island", "Europe/Athens"}
	case "JSI":
		return Location{"Skiathos", "Europe/Athens"}
	case "JSR":
		return Location{"Jashahor", "Asia/Dhaka"}
	case "JST":
		return Location{"Johnstown", "America/New_York"}
	case "JSU":
		return Location{"Maniitsoq", "America/Nuuk"}
	case "JSY":
		return Location{"Syros Island", "Europe/Athens"}
	case "JTC":
		return Location{"Bauru", "America/Sao_Paulo"}
	case "JTR":
		return Location{"Santorini Island", "Europe/Athens"}
	case "JTY":
		return Location{"Astypalaia Island", "Europe/Athens"}
	case "JUB":
		return Location{"Juba", "Africa/Juba"}
	case "JUJ":
		return Location{"San Salvador de Jujuy", "America/Argentina/Jujuy"}
	case "JUL":
		return Location{"Juliaca", "America/Lima"}
	case "JUV":
		return Location{"Upernavik", "America/Nuuk"}
	case "JUZ":
		return Location{"Quzhou", "Asia/Shanghai"}
	case "JYV":
		return Location{"Jyvaskylan Maalaiskunta", "Europe/Helsinki"}
	case "JZH":
		return Location{"Jiuzhaigou", "Asia/Shanghai"}
	case "KAA":
		return Location{"Kasama", "Africa/Lusaka"}
	case "KAB":
		return Location{"Kariba", "Africa/Harare"}
	case "KAD":
		return Location{"Kaduna", "Africa/Lagos"}
	case "KAJ":
		return Location{"Kajaani", "Europe/Helsinki"}
	case "KAL":
		return Location{"Kaltag", "America/Anchorage"}
	case "KAN":
		return Location{"Kano", "Africa/Lagos"}
	case "KAO":
		return Location{"Kuusamo", "Europe/Helsinki"}
	case "KAW":
		return Location{"Kawthoung", "Asia/Yangon"}
	case "KAZ":
		return Location{"Kao-Celebes Island", "Asia/Jayapura"}
	case "KBH":
		return Location{"Kalat", "Asia/Karachi"}
	case "KBL":
		return Location{"Kabul", "Asia/Kabul"}
	case "KBP":
		return Location{"Kiev", "Europe/Kiev"}
	case "KBR":
		return Location{"Kota Baharu", "Asia/Kuala_Lumpur"}
	case "KBU":
		return Location{"Laut Island", "Asia/Makassar"}
	case "KBV":
		return Location{"Krabi", "Asia/Bangkok"}
	case "KCA":
		return Location{"Kuqa", "Asia/Shanghai"}
	case "KCH":
		return Location{"Kuching", "Asia/Kuching"}
	case "KCK":
		return Location{"Kirensk", "Asia/Irkutsk"}
	case "KCM":
		return Location{"Kahramanmaras", "Europe/Istanbul"}
	case "KCT":
		return Location{"Galle", "Asia/Colombo"}
	case "KCZ":
		return Location{"Nankoku", "Asia/Tokyo"}
	case "KDH":
		return Location{"", "Asia/Kabul"}
	case "KDI":
		return Location{"Kendari-Celebes Island", "Asia/Makassar"}
	case "KDM":
		return Location{"Huvadhu Atoll", "Indian/Maldives"}
	case "KDO":
		return Location{"Kadhdhoo", "Indian/Maldives"}
	case "KDU":
		return Location{"Skardu", "Asia/Karachi"}
	case "KDV":
		return Location{"Vunisea", "Pacific/Fiji"}
	case "KEF":
		return Location{"Reykjavik", "Atlantic/Reykjavik"}
	case "KEJ":
		return Location{"Kemerovo", "Asia/Novokuznetsk"}
	case "KEM":
		return Location{"Kemi / Tornio", "Europe/Helsinki"}
	case "KEO":
		return Location{"Odienne", "Africa/Abidjan"}
	case "KEP":
		return Location{"Nepalgunj", "Asia/Kathmandu"}
	case "KER":
		return Location{"Kerman", "Asia/Tehran"}
	case "KET":
		return Location{"Kengtung", "Asia/Yangon"}
	case "KEU":
		return Location{"Keekorok", "Africa/Nairobi"}
	case "KFP":
		return Location{"False Pass", "America/Nome"}
	case "KFS":
		return Location{"Kastamonu", "Europe/Istanbul"}
	case "KGA":
		return Location{"Kananga", "Africa/Lubumbashi"}
	case "KGC":
		return Location{"", "Australia/Adelaide"}
	case "KGD":
		return Location{"Kaliningrad", "Europe/Kaliningrad"}
	case "KGE":
		return Location{"Kagau Island", "Pacific/Guadalcanal"}
	case "KGF":
		return Location{"Karaganda", "Asia/Almaty"}
	case "KGI":
		return Location{"Kalgoorlie", "Australia/Perth"}
	case "KGK":
		return Location{"Koliganek", "America/Anchorage"}
	case "KGL":
		return Location{"Kigali", "Africa/Kigali"}
	case "KGP":
		return Location{"Kogalym", "Asia/Yekaterinburg"}
	case "KGS":
		return Location{"Kos Island", "Europe/Athens"}
	case "KHG":
		return Location{"Kashgar", "Asia/Shanghai"}
	case "KHH":
		return Location{"Kaohsiung City", "Asia/Taipei"}
	case "KHI":
		return Location{"Karachi", "Asia/Karachi"}
	case "KHM":
		return Location{"Kanti", "Asia/Yangon"}
	case "KHN":
		return Location{"Nanchang", "Asia/Shanghai"}
	case "KHS":
		return Location{"Khasab", "Asia/Muscat"}
	case "KHT":
		return Location{"Khost", "Asia/Kabul"}
	case "KHV":
		return Location{"Khabarovsk", "Asia/Vladivostok"}
	case "KID":
		return Location{"Kristianstad", "Europe/Stockholm"}
	case "KIE":
		return Location{"Kieta", "Pacific/Bougainville"}
	case "KIH":
		return Location{"Kish Island", "Asia/Tehran"}
	case "KIJ":
		return Location{"Niigata", "Asia/Tokyo"}
	case "KIK":
		return Location{"Kirkuk", "Asia/Baghdad"}
	case "KIM":
		return Location{"Kimberley", "Africa/Johannesburg"}
	case "KIN":
		return Location{"Kingston", "America/Jamaica"}
	case "KIR":
		return Location{"Killarney", "Europe/Dublin"}
	case "KIS":
		return Location{"Kisumu", "Africa/Nairobi"}
	case "KIT":
		return Location{"Kithira Island", "Europe/Athens"}
	case "KIX":
		return Location{"Osaka", "Asia/Tokyo"}
	case "KJA":
		return Location{"Krasnoyarsk", "Asia/Krasnoyarsk"}
	case "KKA":
		return Location{"Koyuk", "America/Anchorage"}
	case "KKC":
		return Location{"Khon Kaen", "Asia/Bangkok"}
	case "KKE":
		return Location{"Kerikeri", "Pacific/Auckland"}
	case "KKH":
		return Location{"Kongiganak", "America/Nome"}
	case "KKJ":
		return Location{"Kitakyushu", "Asia/Tokyo"}
	case "KKN":
		return Location{"Kirkenes", "Europe/Oslo"}
	case "KKQ":
		return Location{"Krasnoselkup", "Asia/Yekaterinburg"}
	case "KKR":
		return Location{"", "Pacific/Tahiti"}
	case "KKS":
		return Location{"", "Asia/Tehran"}
	case "KKX":
		return Location{"", "Asia/Tokyo"}
	case "KLF":
		return Location{"Kaluga", "Europe/Moscow"}
	case "KLG":
		return Location{"Kalskag", "America/Anchorage"}
	case "KLH":
		return Location{"", "Asia/Kolkata"}
	case "KLN":
		return Location{"Larsen Bay", "America/Anchorage"}
	case "KLO":
		return Location{"Kalibo", "Asia/Manila"}
	case "KLR":
		return Location{"", "Europe/Stockholm"}
	case "KLU":
		return Location{"Klagenfurt am Worthersee", "Europe/Vienna"}
	case "KLX":
		return Location{"Kalamata", "Europe/Athens"}
	case "KME":
		return Location{"Kamembe", "Africa/Kigali"}
	case "KMG":
		return Location{"Kunming", "Asia/Shanghai"}
	case "KMI":
		return Location{"Miyazaki", "Asia/Tokyo"}
	case "KMJ":
		return Location{"Kumamoto", "Asia/Tokyo"}
	case "KMO":
		return Location{"Manokotak", "America/Anchorage"}
	case "KMQ":
		return Location{"Kanazawa", "Asia/Tokyo"}
	case "KMS":
		return Location{"Kumasi", "Africa/Accra"}
	case "KMV":
		return Location{"Kalemyo", "Asia/Yangon"}
	case "KND":
		return Location{"Kindu", "Africa/Lubumbashi"}
	case "KNG":
		return Location{"Kaimana-Papua Island", "Asia/Jayapura"}
	case "KNH":
		return Location{"Shang-I", "Asia/Taipei"}
	case "KNK":
		return Location{"Kokhanok", "America/Anchorage"}
	case "KNO":
		return Location{"Medan-Sumatra Island", "Asia/Jakarta"}
	case "KNQ":
		return Location{"Kone", "Pacific/Noumea"}
	case "KNS":
		return Location{"", "Australia/Currie"}
	case "KNU":
		return Location{"", "Asia/Kolkata"}
	case "KNW":
		return Location{"New Stuyahok", "America/Anchorage"}
	case "KNX":
		return Location{"Kununurra", "Australia/Perth"}
	case "KOA":
		return Location{"Kailua/Kona", "Pacific/Honolulu"}
	case "KOE":
		return Location{"Kupang-Timor Island", "Asia/Makassar"}
	case "KOI":
		return Location{"Orkney Islands", "Europe/London"}
	case "KOJ":
		return Location{"Kagoshima", "Asia/Tokyo"}
	case "KOK":
		return Location{"Kokkola / Kruunupyy", "Europe/Helsinki"}
	case "KOO":
		return Location{"Kongolo", "Africa/Lubumbashi"}
	case "KOP":
		return Location{"", "Asia/Bangkok"}
	case "KOS":
		return Location{"Sihanukville", "Asia/Phnom_Penh"}
	case "KOT":
		return Location{"Kotlik", "America/Nome"}
	case "KOV":
		return Location{"Kokshetau", "Asia/Almaty"}
	case "KOW":
		return Location{"Ganzhou", "Asia/Shanghai"}
	case "KPN":
		return Location{"Kipnuk", "America/Nome"}
	case "KPO":
		return Location{"Pohang", "Asia/Seoul"}
	case "KPV":
		return Location{"Perryville", "America/Anchorage"}
	case "KPW":
		return Location{"Keperveem", "Asia/Anadyr"}
	case "KQH":
		return Location{"Kishangarh", "Asia/Kolkata"}
	case "KQT":
		return Location{"Kurgan-Tyube", "Asia/Dushanbe"}
	case "KRF":
		return Location{"Kramfors / Solleftea", "Europe/Stockholm"}
	case "KRK":
		return Location{"Krakow", "Europe/Warsaw"}
	case "KRL":
		return Location{"Korla", "Asia/Shanghai"}
	case "KRN":
		return Location{"", "Europe/Stockholm"}
	case "KRO":
		return Location{"Kurgan", "Asia/Yekaterinburg"}
	case "KRS":
		return Location{"Kjevik", "Europe/Oslo"}
	case "KRT":
		return Location{"Khartoum", "Africa/Khartoum"}
	case "KRW":
		return Location{"Krasnovodsk", "Asia/Ashgabat"}
	case "KRY":
		return Location{"Karamay", "Asia/Shanghai"}
	case "KSA":
		return Location{"Okat", "Pacific/Kosrae"}
	case "KSC":
		return Location{"Kosice", "Europe/Bratislava"}
	case "KSE":
		return Location{"Kasese", "Africa/Kampala"}
	case "KSF":
		return Location{"Kassel", "Europe/Berlin"}
	case "KSH":
		return Location{"Kermanshah", "Asia/Tehran"}
	case "KSJ":
		return Location{"Kasos Island", "Europe/Athens"}
	case "KSM":
		return Location{"St Mary's", "America/Nome"}
	case "KSN":
		return Location{"Kostanay", "Asia/Qostanay"}
	case "KSO":
		return Location{"Kastoria", "Europe/Athens"}
	case "KSQ":
		return Location{"Karshi", "Asia/Samarkand"}
	case "KSU":
		return Location{"Kvernberget", "Europe/Oslo"}
	case "KSY":
		return Location{"Kars", "Europe/Istanbul"}
	case "KSZ":
		return Location{"Kotlas", "Europe/Moscow"}
	case "KTA":
		return Location{"Karratha", "Australia/Perth"}
	case "KTG":
		return Location{"Ketapang-Borneo Island", "Asia/Pontianak"}
	case "KTL":
		return Location{"Kitale", "Africa/Nairobi"}
	case "KTM":
		return Location{"Kathmandu", "Asia/Kathmandu"}
	case "KTN":
		return Location{"Ketchikan", "America/Sitka"}
	case "KTR":
		return Location{"", "Australia/Darwin"}
	case "KTS":
		return Location{"Brevig Mission", "America/Nome"}
	case "KTT":
		return Location{"Kittila", "Europe/Helsinki"}
	case "KTW":
		return Location{"Katowice", "Europe/Warsaw"}
	case "KUA":
		return Location{"Kuantan", "Asia/Kuala_Lumpur"}
	case "KUF":
		return Location{"Samara", "Europe/Samara"}
	case "KUG":
		return Location{"", "Australia/Brisbane"}
	case "KUH":
		return Location{"Kushiro", "Asia/Tokyo"}
	case "KUK":
		return Location{"Kasigluk", "America/Nome"}
	case "KUL":
		return Location{"Kuala Lumpur", "Asia/Kuala_Lumpur"}
	case "KUM":
		return Location{"", "Asia/Tokyo"}
	case "KUN":
		return Location{"Kaunas", "Europe/Vilnius"}
	case "KUO":
		return Location{"Kuopio / Siilinjarvi", "Europe/Helsinki"}
	case "KUS":
		return Location{"Kulusuk", "America/Nuuk"}
	case "KUT":
		return Location{"Kutaisi", "Asia/Tbilisi"}
	case "KUU":
		return Location{"", "Asia/Kolkata"}
	case "KUV":
		return Location{"Kunsan", "Asia/Seoul"}
	case "KVA":
		return Location{"Kavala", "Europe/Athens"}
	case "KVC":
		return Location{"King Cove", "America/Nome"}
	case "KVG":
		return Location{"Kavieng", "Pacific/Port_Moresby"}
	case "KVK":
		return Location{"Apatity", "Europe/Moscow"}
	case "KVL":
		return Location{"Kivalina", "America/Nome"}
	case "KVX":
		return Location{"Kirov", "Europe/Kirov"}
	case "KWA":
		return Location{"Kwajalein", "Pacific/Kwajalein"}
	case "KWE":
		return Location{"Guiyang", "Asia/Shanghai"}
	case "KWI":
		return Location{"Kuwait City", "Asia/Kuwait"}
	case "KWJ":
		return Location{"Gwangju", "Asia/Seoul"}
	case "KWK":
		return Location{"Kwigillingok", "America/Nome"}
	case "KWL":
		return Location{"Guilin City", "Asia/Shanghai"}
	case "KWM":
		return Location{"Kowanyama", "Australia/Brisbane"}
	case "KWN":
		return Location{"Quinhagak", "America/Anchorage"}
	case "KWT":
		return Location{"Kwethluk", "America/Anchorage"}
	case "KWZ":
		return Location{"", "Africa/Lubumbashi"}
	case "KXB":
		return Location{"Kolaka", "Asia/Makassar"}
	case "KXF":
		return Location{"Koro Island", "Pacific/Fiji"}
	case "KXU":
		return Location{"Katiu", "Pacific/Tahiti"}
	case "KYA":
		return Location{"Konya", "Europe/Istanbul"}
	case "KYK":
		return Location{"Karluk", "America/Anchorage"}
	case "KYP":
		return Location{"Kyaukpyu", "Asia/Yangon"}
	case "KYU":
		return Location{"Koyukuk", "America/Anchorage"}
	case "KYZ":
		return Location{"Kyzyl", "Asia/Krasnoyarsk"}
	case "KZI":
		return Location{"Kozani", "Europe/Athens"}
	case "KZN":
		return Location{"Kazan", "Europe/Moscow"}
	case "KZO":
		return Location{"Kzyl-Orda", "Asia/Qyzylorda"}
	case "KZR":
		return Location{"Altintas", "Europe/Istanbul"}
	case "KZS":
		return Location{"Kastelorizo Island", "Europe/Athens"}
	case "LAD":
		return Location{"Luanda", "Africa/Luanda"}
	case "LAE":
		return Location{"Nadzab", "Pacific/Port_Moresby"}
	case "LAH":
		return Location{"Labuha-Halmahera Island", "Asia/Jayapura"}
	case "LAK":
		return Location{"Aklavik", "America/Inuvik"}
	case "LAN":
		return Location{"Lansing", "America/Detroit"}
	case "LAO":
		return Location{"Laoag City", "Asia/Manila"}
	case "LAP":
		return Location{"La Paz", "America/Mazatlan"}
	case "LAQ":
		return Location{"Al Bayda'", "Africa/Tripoli"}
	case "LAR":
		return Location{"Laramie", "America/Denver"}
	case "LAS":
		return Location{"Las Vegas", "America/Los_Angeles"}
	case "LAU":
		return Location{"Lamu", "Africa/Nairobi"}
	case "LAW":
		return Location{"Lawton", "America/Chicago"}
	case "LAX":
		return Location{"Los Angeles", "America/Los_Angeles"}
	case "LBA":
		return Location{"Leeds", "Europe/London"}
	case "LBB":
		return Location{"Lubbock", "America/Chicago"}
	case "LBD":
		return Location{"Khudzhand", "Asia/Dushanbe"}
	case "LBE":
		return Location{"Latrobe", "America/New_York"}
	case "LBF":
		return Location{"North Platte", "America/Chicago"}
	case "LBJ":
		return Location{"Labuan Bajo-Flores Island", "Asia/Makassar"}
	case "LBL":
		return Location{"Liberal", "America/Chicago"}
	case "LBR":
		return Location{"Labrea", "America/Manaus"}
	case "LBS":
		return Location{"", "Pacific/Fiji"}
	case "LBU":
		return Location{"Labuan", "Asia/Kuching"}
	case "LBV":
		return Location{"Libreville", "Africa/Libreville"}
	case "LCA":
		return Location{"Larnarca", "Asia/Nicosia"}
	case "LCE":
		return Location{"La Ceiba", "America/Tegucigalpa"}
	case "LCG":
		return Location{"Culleredo", "Europe/Madrid"}
	case "LCH":
		return Location{"Lake Charles", "America/Chicago"}
	case "LCJ":
		return Location{"Lodz", "Europe/Warsaw"}
	case "LCK":
		return Location{"Columbus", "America/New_York"}
	case "LCX":
		return Location{"Longyan", "Asia/Shanghai"}
	case "LCY":
		return Location{"London", "Europe/London"}
	case "LDB":
		return Location{"Londrina", "America/Sao_Paulo"}
	case "LDE":
		return Location{"Tarbes/Lourdes/Pyrenees", "Europe/Paris"}
	case "LDH":
		return Location{"Lord Howe Island", "Australia/Lord_Howe"}
	case "LDS":
		return Location{"Yichun", "Asia/Shanghai"}
	case "LDU":
		return Location{"Lahad Datu", "Asia/Kuching"}
	case "LDY":
		return Location{"Derry", "Europe/London"}
	case "LEA":
		return Location{"Exmouth", "Australia/Perth"}
	case "LEB":
		return Location{"Lebanon", "America/New_York"}
	case "LEC":
		return Location{"Lencois", "America/Bahia"}
	case "LED":
		return Location{"St. Petersburg", "Europe/Moscow"}
	case "LEI":
		return Location{"Almeria", "Europe/Madrid"}
	case "LEJ":
		return Location{"Leipzig", "Europe/Berlin"}
	case "LEN":
		return Location{"Leon", "Europe/Madrid"}
	case "LET":
		return Location{"Leticia", "America/Bogota"}
	case "LEU":
		return Location{"Montferrer / Castellbo", "Europe/Madrid"}
	case "LEX":
		return Location{"Lexington", "America/New_York"}
	case "LFR":
		return Location{"", "America/Caracas"}
	case "LFT":
		return Location{"Lafayette", "America/Chicago"}
	case "LFW":
		return Location{"Lome", "Africa/Lome"}
	case "LGA":
		return Location{"New York", "America/New_York"}
	case "LGB":
		return Location{"Long Beach", "America/Los_Angeles"}
	case "LGG":
		return Location{"Liege", "Europe/Brussels"}
	case "LGI":
		return Location{"Deadman's Cay", "America/Nassau"}
	case "LGK":
		return Location{"Langkawi", "Asia/Kuala_Lumpur"}
	case "LGL":
		return Location{"Long Datih", "Asia/Kuching"}
	case "LGW":
		return Location{"London", "Europe/London"}
	case "LHE":
		return Location{"Lahore", "Asia/Karachi"}
	case "LHG":
		return Location{"", "Australia/Sydney"}
	case "LHR":
		return Location{"London", "Europe/London"}
	case "LHW":
		return Location{"Lanzhou", "Asia/Shanghai"}
	case "LIF":
		return Location{"Lifou", "Pacific/Noumea"}
	case "LIG":
		return Location{"Limoges/Bellegarde", "Europe/Paris"}
	case "LIH":
		return Location{"Lihue", "Pacific/Honolulu"}
	case "LIL":
		return Location{"Lille/Lesquin", "Europe/Paris"}
	case "LIM":
		return Location{"Lima", "America/Lima"}
	case "LIN":
		return Location{"Milan", "Europe/Rome"}
	case "LIO":
		return Location{"Puerto Limon", "America/Costa_Rica"}
	case "LIR":
		return Location{"Liberia", "America/Costa_Rica"}
	case "LIS":
		return Location{"Lisbon", "Europe/Lisbon"}
	case "LIT":
		return Location{"Little Rock", "America/Chicago"}
	case "LJG":
		return Location{"Lijiang", "Asia/Shanghai"}
	case "LJU":
		return Location{"Ljubljana", "Europe/Ljubljana"}
	case "LKA":
		return Location{"Larantuka-Flores Island", "Asia/Makassar"}
	case "LKB":
		return Location{"Lakeba Island", "Pacific/Fiji"}
	case "LKH":
		return Location{"Long Akah", "Asia/Kuching"}
	case "LKL":
		return Location{"Lakselv", "Europe/Oslo"}
	case "LKN":
		return Location{"Leknes", "Europe/Oslo"}
	case "LKO":
		return Location{"Lucknow", "Asia/Kolkata"}
	case "LKY":
		return Location{"Lake Manyara National Park", "Africa/Dar_es_Salaam"}
	case "LLA":
		return Location{"Lulea", "Europe/Stockholm"}
	case "LLF":
		return Location{"Yongzhou", "Asia/Shanghai"}
	case "LLI":
		return Location{"Lalibela", "Africa/Addis_Ababa"}
	case "LLJ":
		return Location{"Lalmonirhat", "Asia/Dhaka"}
	case "LLK":
		return Location{"Lankaran", "Asia/Baku"}
	case "LLW":
		return Location{"Lilongwe", "Africa/Blantyre"}
	case "LMA":
		return Location{"Minchumina", "America/Anchorage"}
	case "LMM":
		return Location{"Los Mochis", "America/Mazatlan"}
	case "LMN":
		return Location{"Limbang", "Asia/Brunei"}
	case "LMP":
		return Location{"Lampedusa", "Europe/Rome"}
	case "LNB":
		return Location{"Lamen Bay", "Pacific/Efate"}
	case "LNE":
		return Location{"Lonorore", "Pacific/Efate"}
	case "LNJ":
		return Location{"Lincang", "Asia/Shanghai"}
	case "LNK":
		return Location{"Lincoln", "America/Chicago"}
	case "LNO":
		return Location{"Leonora", "Australia/Perth"}
	case "LNS":
		return Location{"Lancaster", "America/New_York"}
	case "LNV":
		return Location{"Londolovit", "Pacific/Port_Moresby"}
	case "LNY":
		return Location{"Lanai City", "Pacific/Honolulu"}
	case "LNZ":
		return Location{"Linz", "Europe/Vienna"}
	case "LOD":
		return Location{"Longana", "Pacific/Efate"}
	case "LOE":
		return Location{"", "Asia/Bangkok"}
	case "LOH":
		return Location{"La Toma (Catamayo)", "America/Guayaquil"}
	case "LOK":
		return Location{"Lodwar", "Africa/Nairobi"}
	case "LOO":
		return Location{"Laghouat", "Africa/Algiers"}
	case "LOP":
		return Location{"Mataram", "Asia/Makassar"}
	case "LOS":
		return Location{"Lagos", "Africa/Lagos"}
	case "LOY":
		return Location{"Loyengalani", "Africa/Nairobi"}
	case "LPA":
		return Location{"Gran Canaria Island", "Atlantic/Canary"}
	case "LPB":
		return Location{"La Paz / El Alto", "America/La_Paz"}
	case "LPI":
		return Location{"Linkoping", "Europe/Stockholm"}
	case "LPL":
		return Location{"Liverpool", "Europe/London"}
	case "LPM":
		return Location{"Lamap", "Pacific/Efate"}
	case "LPP":
		return Location{"Lappeenranta", "Europe/Helsinki"}
	case "LPQ":
		return Location{"Luang Phabang", "Asia/Vientiane"}
	case "LPT":
		return Location{"", "Asia/Bangkok"}
	case "LPY":
		return Location{"Le Puy/Loudes", "Europe/Paris"}
	case "LQM":
		return Location{"Puerto Leguizamo", "America/Bogota"}
	case "LRD":
		return Location{"Laredo", "America/Chicago"}
	case "LRE":
		return Location{"Longreach", "Australia/Brisbane"}
	case "LRH":
		return Location{"La Rochelle/Ile de Re", "Europe/Paris"}
	case "LRM":
		return Location{"La Romana", "America/Santo_Domingo"}
	case "LRR":
		return Location{"Lar", "Asia/Tehran"}
	case "LRS":
		return Location{"Leros Island", "Europe/Athens"}
	case "LRT":
		return Location{"Lorient/Lann/Bihoue", "Europe/Paris"}
	case "LRU":
		return Location{"Las Cruces", "America/Denver"}
	case "LRV":
		return Location{"Los Roques", "America/Caracas"}
	case "LSC":
		return Location{"La Serena-Coquimbo", "America/Santiago"}
	case "LSE":
		return Location{"La Crosse", "America/Chicago"}
	case "LSH":
		return Location{"Lashio", "Asia/Yangon"}
	case "LSI":
		return Location{"Lerwick", "Europe/London"}
	case "LSP":
		return Location{"Paraguana", "America/Caracas"}
	case "LSQ":
		return Location{"Los Angeles", "America/Santiago"}
	case "LST":
		return Location{"Launceston", "Australia/Hobart"}
	case "LSW":
		return Location{"Lhok Seumawe-Sumatra Island", "Asia/Jakarta"}
	case "LTI":
		return Location{"Altai", "Asia/Hovd"}
	case "LTN":
		return Location{"London", "Europe/London"}
	case "LTO":
		return Location{"Loreto", "America/Mazatlan"}
	case "LUD":
		return Location{"Luderitz", "Africa/Windhoek"}
	case "LUM":
		return Location{"Luxi", "Asia/Shanghai"}
	case "LUN":
		return Location{"Lusaka", "Africa/Lusaka"}
	case "LUO":
		return Location{"Luena", "Africa/Luanda"}
	case "LUP":
		return Location{"Kalaupapa", "Pacific/Honolulu"}
	case "LUQ":
		return Location{"San Luis", "America/Argentina/San_Luis"}
	case "LUR":
		return Location{"Cape Lisburne", "America/Nome"}
	case "LUV":
		return Location{"Langgur-Seram Island", "Asia/Jayapura"}
	case "LUW":
		return Location{"Luwok-Celebes Island", "Asia/Makassar"}
	case "LUX":
		return Location{"Luxembourg", "Europe/Luxembourg"}
	case "LUZ":
		return Location{"Lublin", "Europe/Warsaw"}
	case "LVI":
		return Location{"Livingstone", "Africa/Lusaka"}
	case "LVO":
		return Location{"", "Australia/Perth"}
	case "LWB":
		return Location{"Lewisburg", "America/New_York"}
	case "LWS":
		return Location{"Lewiston", "America/Los_Angeles"}
	case "LWY":
		return Location{"Lawas", "Asia/Kuching"}
	case "LXA":
		return Location{"Lhasa", "Asia/Shanghai"}
	case "LXG":
		return Location{"Luang Namtha", "Asia/Vientiane"}
	case "LXR":
		return Location{"Luxor", "Africa/Cairo"}
	case "LXS":
		return Location{"Limnos Island", "Europe/Athens"}
	case "LYA":
		return Location{"Luoyang", "Asia/Shanghai"}
	case "LYB":
		return Location{"Little Cayman", "America/Cayman"}
	case "LYC":
		return Location{"", "Europe/Stockholm"}
	case "LYG":
		return Location{"Lianyungang", "Asia/Shanghai"}
	case "LYH":
		return Location{"Lynchburg", "America/New_York"}
	case "LYI":
		return Location{"Linyi", "Asia/Shanghai"}
	case "LYP":
		return Location{"Faisalabad", "Asia/Karachi"}
	case "LYR":
		return Location{"Longyearbyen", "Arctic/Longyearbyen"}
	case "LYS":
		return Location{"Lyon", "Europe/Paris"}
	case "LZH":
		return Location{"Liuzhou", "Asia/Shanghai"}
	case "LZN":
		return Location{"Nangang Island", "Asia/Taipei"}
	case "LZO":
		return Location{"Luzhou", "Asia/Shanghai"}
	case "LZY":
		return Location{"Nyingchi", "Asia/Shanghai"}
	case "MAA":
		return Location{"Chennai", "Asia/Kolkata"}
	case "MAB":
		return Location{"Maraba", "America/Belem"}
	case "MAD":
		return Location{"Madrid", "Europe/Madrid"}
	case "MAF":
		return Location{"Midland", "America/Chicago"}
	case "MAG":
		return Location{"Madang", "Pacific/Port_Moresby"}
	case "MAH":
		return Location{"Menorca Island", "Europe/Madrid"}
	case "MAJ":
		return Location{"Majuro Atoll", "Pacific/Majuro"}
	case "MAM":
		return Location{"Matamoros", "America/Matamoros"}
	case "MAN":
		return Location{"Manchester", "Europe/London"}
	case "MAO":
		return Location{"Manaus", "America/Manaus"}
	case "MAQ":
		return Location{"", "Asia/Bangkok"}
	case "MAR":
		return Location{"Maracaibo", "America/Caracas"}
	case "MAS":
		return Location{"", "Pacific/Port_Moresby"}
	case "MAU":
		return Location{"", "Pacific/Tahiti"}
	case "MAZ":
		return Location{"Mayaguez", "America/Puerto_Rico"}
	case "MBA":
		return Location{"Mombasa", "Africa/Nairobi"}
	case "MBE":
		return Location{"Monbetsu", "Asia/Tokyo"}
	case "MBI":
		return Location{"Mbeya", "Africa/Dar_es_Salaam"}
	case "MBJ":
		return Location{"Montego Bay", "America/Jamaica"}
	case "MBL":
		return Location{"Manistee", "America/Detroit"}
	case "MBQ":
		return Location{"Mbarara", "Africa/Kampala"}
	case "MBS":
		return Location{"Saginaw", "America/Detroit"}
	case "MBT":
		return Location{"Masbate", "Asia/Manila"}
	case "MBZ":
		return Location{"Maues", "America/Manaus"}
	case "MCE":
		return Location{"Merced", "America/Los_Angeles"}
	case "MCI":
		return Location{"Kansas City", "America/Chicago"}
	case "MCK":
		return Location{"Mc Cook", "America/Chicago"}
	case "MCN":
		return Location{"Macon", "America/New_York"}
	case "MCO":
		return Location{"Orlando", "America/New_York"}
	case "MCP":
		return Location{"Macapa", "America/Belem"}
	case "MCT":
		return Location{"Muscat", "Asia/Muscat"}
	case "MCV":
		return Location{"McArthur River Mine", "Australia/Darwin"}
	case "MCW":
		return Location{"Mason City", "America/Chicago"}
	case "MCX":
		return Location{"Makhachkala", "Europe/Moscow"}
	case "MCY":
		return Location{"Maroochydore", "Australia/Brisbane"}
	case "MCZ":
		return Location{"Maceio", "America/Maceio"}
	case "MDC":
		return Location{"Manado-Celebes Island", "Asia/Makassar"}
	case "MDE":
		return Location{"Rionegro", "America/Bogota"}
	case "MDG":
		return Location{"Mudanjiang", "Asia/Shanghai"}
	case "MDI":
		return Location{"Makurdi", "Africa/Lagos"}
	case "MDK":
		return Location{"Mbandaka", "Africa/Kinshasa"}
	case "MDL":
		return Location{"Mandalay", "Asia/Yangon"}
	case "MDQ":
		return Location{"Mar del Plata", "America/Argentina/Buenos_Aires"}
	case "MDT":
		return Location{"Harrisburg", "America/New_York"}
	case "MDW":
		return Location{"Chicago", "America/Chicago"}
	case "MDZ":
		return Location{"Mendoza", "America/Argentina/Mendoza"}
	case "MEB":
		return Location{"", "Australia/Melbourne"}
	case "MEC":
		return Location{"Manta", "America/Guayaquil"}
	case "MED":
		return Location{"Medina", "Asia/Riyadh"}
	case "MEE":
		return Location{"Mare", "Pacific/Noumea"}
	case "MEH":
		return Location{"Mehamn", "Europe/Oslo"}
	case "MEI":
		return Location{"Meridian", "America/Chicago"}
	case "MEL":
		return Location{"Melbourne", "Australia/Melbourne"}
	case "MEM":
		return Location{"Memphis", "America/Chicago"}
	case "MEQ":
		return Location{"Peureumeue-Sumatra Island", "Asia/Jakarta"}
	case "MEU":
		return Location{"Almeirim", "America/Santarem"}
	case "MEX":
		return Location{"Mexico City", "America/Mexico_City"}
	case "MFA":
		return Location{"Mafia Island", "Africa/Dar_es_Salaam"}
	case "MFE":
		return Location{"Mc Allen", "America/Chicago"}
	case "MFK":
		return Location{"Beigan Island", "Asia/Taipei"}
	case "MFM":
		return Location{"Taipa", "Asia/Macau"}
	case "MFR":
		return Location{"Medford", "America/Los_Angeles"}
	case "MFU":
		return Location{"Mfuwe", "Africa/Lusaka"}
	case "MGA":
		return Location{"Managua", "America/Managua"}
	case "MGB":
		return Location{"", "Australia/Adelaide"}
	case "MGF":
		return Location{"Maringa", "America/Sao_Paulo"}
	case "MGH":
		return Location{"Margate", "Africa/Johannesburg"}
	case "MGM":
		return Location{"Montgomery", "America/Chicago"}
	case "MGQ":
		return Location{"Mogadishu", "Africa/Mogadishu"}
	case "MGT":
		return Location{"Milingimbi Island", "Australia/Darwin"}
	case "MGW":
		return Location{"Morgantown", "America/New_York"}
	case "MGZ":
		return Location{"Mkeik", "Asia/Yangon"}
	case "MHC":
		return Location{"Dalcahue", "America/Santiago"}
	case "MHD":
		return Location{"Mashhad", "Asia/Tehran"}
	case "MHH":
		return Location{"Marsh Harbour", "America/Nassau"}
	case "MHK":
		return Location{"Manhattan", "America/Chicago"}
	case "MHQ":
		return Location{"", "Europe/Mariehamn"}
	case "MHT":
		return Location{"Manchester", "America/New_York"}
	case "MIA":
		return Location{"Miami", "America/New_York"}
	case "MID":
		return Location{"Merida", "America/Merida"}
	case "MIG":
		return Location{"Mianyang", "Asia/Shanghai"}
	case "MII":
		return Location{"Marilia", "America/Sao_Paulo"}
	case "MIM":
		return Location{"Merimbula", "Australia/Sydney"}
	case "MIR":
		return Location{"Monastir", "Africa/Tunis"}
	case "MIS":
		return Location{"Misima Island", "Pacific/Port_Moresby"}
	case "MIU":
		return Location{"Maiduguri", "Africa/Lagos"}
	case "MJC":
		return Location{"", "Africa/Abidjan"}
	case "MJF":
		return Location{"", "Europe/Oslo"}
	case "MJI":
		return Location{"Tripoli", "Africa/Tripoli"}
	case "MJK":
		return Location{"Monkey Mia", "Australia/Perth"}
	case "MJM":
		return Location{"Mbuji Mayi", "Africa/Lubumbashi"}
	case "MJN":
		return Location{"", "Indian/Antananarivo"}
	case "MJT":
		return Location{"Mytilene", "Europe/Athens"}
	case "MJU":
		return Location{"Mamuju-Celebes Island", "Asia/Makassar"}
	case "MJZ":
		return Location{"Mirny", "Asia/Yakutsk"}
	case "MKE":
		return Location{"Milwaukee", "America/Chicago"}
	case "MKG":
		return Location{"Muskegon", "America/Detroit"}
	case "MKK":
		return Location{"Kaunakakai", "Pacific/Honolulu"}
	case "MKL":
		return Location{"Jackson", "America/Chicago"}
	case "MKM":
		return Location{"Mukah", "Asia/Kuching"}
	case "MKP":
		return Location{"", "Pacific/Tahiti"}
	case "MKQ":
		return Location{"Merauke-Papua Island", "Asia/Jayapura"}
	case "MKR":
		return Location{"", "Australia/Perth"}
	case "MKW":
		return Location{"Manokwari-Papua Island", "Asia/Jayapura"}
	case "MKY":
		return Location{"Mackay", "Australia/Brisbane"}
	case "MLA":
		return Location{"Luqa", "Europe/Malta"}
	case "MLB":
		return Location{"Melbourne", "America/New_York"}
	case "MLE":
		return Location{"Male", "Indian/Maldives"}
	case "MLG":
		return Location{"Malang-Java Island", "Asia/Jakarta"}
	case "MLH":
		return Location{"Bale/Mulhouse", "Europe/Paris"}
	case "MLI":
		return Location{"Moline", "America/Chicago"}
	case "MLL":
		return Location{"Marshall", "America/Anchorage"}
	case "MLM":
		return Location{"Morelia", "America/Mexico_City"}
	case "MLN":
		return Location{"Melilla Island", "Africa/Casablanca"}
	case "MLO":
		return Location{"Milos Island", "Europe/Athens"}
	case "MLU":
		return Location{"Monroe", "America/Chicago"}
	case "MLX":
		return Location{"Malatya", "Europe/Istanbul"}
	case "MLY":
		return Location{"Manley Hot Springs", "America/Anchorage"}
	case "MMB":
		return Location{"Ozora", "Asia/Tokyo"}
	case "MME":
		return Location{"Durham", "Europe/London"}
	case "MMG":
		return Location{"", "Australia/Perth"}
	case "MMH":
		return Location{"Mammoth Lakes", "America/Los_Angeles"}
	case "MMJ":
		return Location{"Matsumoto", "Asia/Tokyo"}
	case "MMK":
		return Location{"Murmansk", "Europe/Moscow"}
	case "MMO":
		return Location{"Vila do Maio", "Atlantic/Cape_Verde"}
	case "MMU":
		return Location{"Morristown", "America/New_York"}
	case "MMX":
		return Location{"Malmo", "Europe/Stockholm"}
	case "MMY":
		return Location{"Miyako City", "Asia/Tokyo"}
	case "MNA":
		return Location{"Karakelong Island", "Asia/Makassar"}
	case "MNC":
		return Location{"Nacala", "Africa/Maputo"}
	case "MNG":
		return Location{"Maningrida", "Australia/Darwin"}
	case "MNI":
		return Location{"Gerald's Park", "America/Montserrat"}
	case "MNL":
		return Location{"Manila", "Asia/Manila"}
	case "MNS":
		return Location{"Mansa", "Africa/Lusaka"}
	case "MNX":
		return Location{"Manicore", "America/Manaus"}
	case "MOB":
		return Location{"Mobile", "America/Chicago"}
	case "MOC":
		return Location{"Montes Claros", "America/Sao_Paulo"}
	case "MOF":
		return Location{"Maumere-Flores Island", "Asia/Makassar"}
	case "MOL":
		return Location{"Aro", "Europe/Oslo"}
	case "MOQ":
		return Location{"", "Indian/Antananarivo"}
	case "MOT":
		return Location{"Minot", "America/Chicago"}
	case "MOU":
		return Location{"Mountain Village", "America/Nome"}
	case "MOV":
		return Location{"Moranbah", "Australia/Brisbane"}
	case "MOZ":
		return Location{"", "Pacific/Tahiti"}
	case "MPA":
		return Location{"Mpacha", "Africa/Windhoek"}
	case "MPH":
		return Location{"Malay", "Asia/Manila"}
	case "MPL":
		return Location{"Montpellier/Mediterranee", "Europe/Paris"}
	case "MPM":
		return Location{"Maputo", "Africa/Maputo"}
	case "MPN":
		return Location{"Mount Pleasant", "Atlantic/Stanley"}
	case "MQF":
		return Location{"Magnitogorsk", "Asia/Yekaterinburg"}
	case "MQJ":
		return Location{"Honuu", "Asia/Srednekolymsk"}
	case "MQL":
		return Location{"Mildura", "Australia/Melbourne"}
	case "MQM":
		return Location{"Mardin", "Europe/Istanbul"}
	case "MQN":
		return Location{"Mo i Rana", "Europe/Oslo"}
	case "MQP":
		return Location{"Mpumalanga", "Africa/Johannesburg"}
	case "MQT":
		return Location{"Marquette", "America/Detroit"}
	case "MQX":
		return Location{"", "Africa/Addis_Ababa"}
	case "MRD":
		return Location{"Merida", "America/Caracas"}
	case "MRE":
		return Location{"Masai Mara", "Africa/Nairobi"}
	case "MRS":
		return Location{"Marseille", "Europe/Paris"}
	case "MRU":
		return Location{"Port Louis", "Indian/Mauritius"}
	case "MRV":
		return Location{"Mineralnyye Vody", "Europe/Moscow"}
	case "MRX":
		return Location{"", "Asia/Tehran"}
	case "MRY":
		return Location{"Monterey", "America/Los_Angeles"}
	case "MRZ":
		return Location{"Moree", "Australia/Sydney"}
	case "MSA":
		return Location{"Muskrat Dam", "America/Rainy_River"}
	case "MSH":
		return Location{"Masirah", "Asia/Muscat"}
	case "MSJ":
		return Location{"Misawa", "Asia/Tokyo"}
	case "MSL":
		return Location{"Muscle Shoals", "America/Chicago"}
	case "MSN":
		return Location{"Madison", "America/Chicago"}
	case "MSO":
		return Location{"Missoula", "America/Denver"}
	case "MSP":
		return Location{"Minneapolis", "America/Chicago"}
	case "MSQ":
		return Location{"Minsk", "Europe/Minsk"}
	case "MSR":
		return Location{"Mus", "Europe/Istanbul"}
	case "MSS":
		return Location{"Massena", "America/New_York"}
	case "MST":
		return Location{"Maastricht", "Europe/Amsterdam"}
	case "MSU":
		return Location{"Maseru", "Africa/Maseru"}
	case "MSY":
		return Location{"New Orleans", "America/Chicago"}
	case "MSZ":
		return Location{"Namibe", "Africa/Luanda"}
	case "MTE":
		return Location{"Monte Alegre", "America/Santarem"}
	case "MTJ":
		return Location{"Montrose", "America/Denver"}
	case "MTR":
		return Location{"Monteria", "America/Bogota"}
	case "MTT":
		return Location{"Minatitlan", "America/Mexico_City"}
	case "MTV":
		return Location{"Ablow", "Pacific/Efate"}
	case "MTY":
		return Location{"Monterrey", "America/Monterrey"}
	case "MUA":
		return Location{"", "Pacific/Guadalcanal"}
	case "MUB":
		return Location{"Maun", "Africa/Gaborone"}
	case "MUC":
		return Location{"Munich", "Europe/Berlin"}
	case "MUE":
		return Location{"Kamuela", "Pacific/Honolulu"}
	case "MUH":
		return Location{"Mersa Matruh", "Africa/Cairo"}
	case "MUN":
		return Location{"", "America/Caracas"}
	case "MUR":
		return Location{"Marudi", "Asia/Kuching"}
	case "MUX":
		return Location{"Multan", "Asia/Karachi"}
	case "MVB":
		return Location{"Franceville", "Africa/Libreville"}
	case "MVD":
		return Location{"Montevideo", "America/Montevideo"}
	case "MVF":
		return Location{"Mossoro", "America/Fortaleza"}
	case "MVP":
		return Location{"Mitu", "America/Bogota"}
	case "MVR":
		return Location{"Maroua", "Africa/Douala"}
	case "MVT":
		return Location{"", "Pacific/Tahiti"}
	case "MVY":
		return Location{"Martha's Vineyard", "America/New_York"}
	case "MWA":
		return Location{"Marion", "America/Chicago"}
	case "MWF":
		return Location{"Maewo Island", "Pacific/Efate"}
	case "MWX":
		return Location{"", "Asia/Seoul"}
	case "MWZ":
		return Location{"Mwanza", "Africa/Dar_es_Salaam"}
	case "MXH":
		return Location{"Moro", "Pacific/Port_Moresby"}
	case "MXL":
		return Location{"Mexicali", "America/Tijuana"}
	case "MXP":
		return Location{"Milan", "Europe/Rome"}
	case "MXV":
		return Location{"Moron", "Asia/Ulaanbaatar"}
	case "MXX":
		return Location{"", "Europe/Stockholm"}
	case "MXZ":
		return Location{"Meixian", "Asia/Shanghai"}
	case "MYA":
		return Location{"Moruya", "Australia/Sydney"}
	case "MYD":
		return Location{"Malindi", "Africa/Nairobi"}
	case "MYG":
		return Location{"Mayaguana", "America/Nassau"}
	case "MYI":
		return Location{"Murray Island", "Australia/Brisbane"}
	case "MYJ":
		return Location{"Matsuyama", "Asia/Tokyo"}
	case "MYP":
		return Location{"Mary", "Asia/Ashgabat"}
	case "MYQ":
		return Location{"Mysore", "Asia/Kolkata"}
	case "MYR":
		return Location{"Myrtle Beach", "America/New_York"}
	case "MYT":
		return Location{"Myitkyina", "Asia/Yangon"}
	case "MYU":
		return Location{"Mekoryuk", "America/Nome"}
	case "MYW":
		return Location{"Mtwara", "Africa/Dar_es_Salaam"}
	case "MYY":
		return Location{"Miri", "Asia/Kuching"}
	case "MZA":
		return Location{"Mazamari", "America/Lima"}
	case "MZG":
		return Location{"Makung City", "Asia/Taipei"}
	case "MZH":
		return Location{"Amasya", "Europe/Istanbul"}
	case "MZL":
		return Location{"Manizales", "America/Bogota"}
	case "MZO":
		return Location{"Manzanillo", "America/Havana"}
	case "MZR":
		return Location{"", "Asia/Kabul"}
	case "MZT":
		return Location{"Mazatlan", "America/Mazatlan"}
	case "MZV":
		return Location{"Mulu", "Asia/Kuching"}
	case "MZW":
		return Location{"Mecheria", "Africa/Algiers"}
	case "NAA":
		return Location{"Narrabri", "Australia/Sydney"}
	case "NAG":
		return Location{"Naqpur", "Asia/Kolkata"}
	case "NAH":
		return Location{"Tahuna-Sangihe Island", "Asia/Makassar"}
	case "NAJ":
		return Location{"Nakhchivan", "Asia/Baku"}
	case "NAL":
		return Location{"Nalchik", "Europe/Moscow"}
	case "NAM":
		return Location{"Namlea-Buru Island", "Asia/Jayapura"}
	case "NAN":
		return Location{"Nadi", "Pacific/Fiji"}
	case "NAO":
		return Location{"Nanchong", "Asia/Shanghai"}
	case "NAP":
		return Location{"Napoli", "Europe/Rome"}
	case "NAQ":
		return Location{"Qaanaaq", "America/Thule"}
	case "NAS":
		return Location{"Nassau", "America/Nassau"}
	case "NAT":
		return Location{"Natal", "America/Fortaleza"}
	case "NAU":
		return Location{"Napuka Island", "Pacific/Tahiti"}
	case "NAV":
		return Location{"Nevsehir", "Europe/Istanbul"}
	case "NAW":
		return Location{"", "Asia/Bangkok"}
	case "NBC":
		return Location{"Nizhnekamsk", "Europe/Moscow"}
	case "NBE":
		return Location{"Enfidha", "Africa/Tunis"}
	case "NBO":
		return Location{"Nairobi", "Africa/Nairobi"}
	case "NBX":
		return Location{"Nabire-Papua Island", "Asia/Jayapura"}
	case "NCE":
		return Location{"Nice", "Europe/Paris"}
	case "NCL":
		return Location{"Newcastle", "Europe/London"}
	case "NCU":
		return Location{"Nukus", "Asia/Samarkand"}
	case "NDB":
		return Location{"Nouadhibou", "Africa/El_Aaiun"}
	case "NDG":
		return Location{"Qiqihar", "Asia/Shanghai"}
	case "NDJ":
		return Location{"N'Djamena", "Africa/Ndjamena"}
	case "NDR":
		return Location{"Nador", "Africa/Casablanca"}
	case "NDU":
		return Location{"Rundu", "Africa/Windhoek"}
	case "NEU":
		return Location{"", "Asia/Vientiane"}
	case "NEV":
		return Location{"Charlestown", "America/St_Kitts"}
	case "NGB":
		return Location{"Ningbo", "Asia/Shanghai"}
	case "NGE":
		return Location{"N'Gaoundere", "Africa/Douala"}
	case "NGO":
		return Location{"Tokoname", "Asia/Tokyo"}
	case "NGQ":
		return Location{"Shiquanhe", "Asia/Shanghai"}
	case "NGS":
		return Location{"Nagasaki", "Asia/Tokyo"}
	case "NHV":
		return Location{"", "Pacific/Marquesas"}
	case "NIM":
		return Location{"Niamey", "Africa/Niamey"}
	case "NJC":
		return Location{"Nizhnevartovsk", "Asia/Yekaterinburg"}
	case "NJF":
		return Location{"Najaf", "Asia/Baghdad"}
	case "NKC":
		return Location{"Nouakchott", "Africa/Nouakchott"}
	case "NKG":
		return Location{"Nanjing", "Asia/Shanghai"}
	case "NKM":
		return Location{"Nagoya", "Asia/Tokyo"}
	case "NLA":
		return Location{"Ndola", "Africa/Lusaka"}
	case "NLD":
		return Location{"Nuevo Laredo", "America/Matamoros"}
	case "NLF":
		return Location{"Darnley Island", "Australia/Brisbane"}
	case "NLG":
		return Location{"Nelson Lagoon", "America/Anchorage"}
	case "NLI":
		return Location{"Nikolayevsk-na-Amure Airport", "Asia/Vladivostok"}
	case "NLK":
		return Location{"Burnt Pine", "Pacific/Norfolk"}
	case "NLU":
		return Location{"Santa Lucia", "America/Mexico_City"}
	case "NMA":
		return Location{"Namangan", "Asia/Tashkent"}
	case "NME":
		return Location{"Nightmute", "America/Nome"}
	case "NNB":
		return Location{"Santa Ana Island", "Pacific/Guadalcanal"}
	case "NNG":
		return Location{"Nanning", "Asia/Shanghai"}
	case "NNL":
		return Location{"Nondalton", "America/Anchorage"}
	case "NNM":
		return Location{"Naryan Mar", "Europe/Moscow"}
	case "NNT":
		return Location{"", "Asia/Bangkok"}
	case "NNY":
		return Location{"Nanyang", "Asia/Shanghai"}
	case "NOB":
		return Location{"Nicoya", "America/Costa_Rica"}
	case "NOC":
		return Location{"Charleston", "Europe/Dublin"}
	case "NOJ":
		return Location{"Noyabrsk", "Asia/Yekaterinburg"}
	case "NOP":
		return Location{"Sinop", "Europe/Istanbul"}
	case "NOS":
		return Location{"Nosy Be", "Indian/Antananarivo"}
	case "NOU":
		return Location{"Noumea", "Pacific/Noumea"}
	case "NOV":
		return Location{"Huambo", "Africa/Luanda"}
	case "NOZ":
		return Location{"Novokuznetsk", "Asia/Novokuznetsk"}
	case "NPE":
		return Location{"", "Pacific/Auckland"}
	case "NPL":
		return Location{"New Plymouth", "Pacific/Auckland"}
	case "NQN":
		return Location{"Neuquen", "America/Argentina/Salta"}
	case "NQU":
		return Location{"Nuqui", "America/Bogota"}
	case "NQY":
		return Location{"Newquay", "Europe/London"}
	case "NQZ":
		return Location{"Astana", "Asia/Almaty"}
	case "NRA":
		return Location{"Narrandera", "Australia/Sydney"}
	case "NRK":
		return Location{"Norrkoping", "Europe/Stockholm"}
	case "NRN":
		return Location{"Weeze", "Europe/Amsterdam"}
	case "NRR":
		return Location{"Ceiba", "America/Puerto_Rico"}
	case "NRT":
		return Location{"Tokyo", "Asia/Tokyo"}
	case "NSH":
		return Location{"", "Asia/Tehran"}
	case "NSI":
		return Location{"Yaounde", "Africa/Douala"}
	case "NSK":
		return Location{"Norilsk", "Asia/Krasnoyarsk"}
	case "NSN":
		return Location{"Nelson", "Pacific/Auckland"}
	case "NST":
		return Location{"Nakhon Si Thammarat", "Asia/Bangkok"}
	case "NTE":
		return Location{"Nantes", "Europe/Paris"}
	case "NTG":
		return Location{"Nantong", "Asia/Shanghai"}
	case "NTL":
		return Location{"Williamtown", "Australia/Sydney"}
	case "NTN":
		return Location{"", "Australia/Brisbane"}
	case "NTQ":
		return Location{"Wajima", "Asia/Tokyo"}
	case "NTX":
		return Location{"Ranai-Natuna Besar Island", "Asia/Jakarta"}
	case "NUE":
		return Location{"Nuremberg", "Europe/Berlin"}
	case "NUI":
		return Location{"Nuiqsut", "America/Anchorage"}
	case "NUK":
		return Location{"Nukutavake", "Pacific/Tahiti"}
	case "NUL":
		return Location{"Nulato", "America/Anchorage"}
	case "NUS":
		return Location{"Norsup", "Pacific/Efate"}
	case "NUX":
		return Location{"Novy Urengoy", "Asia/Yekaterinburg"}
	case "NVA":
		return Location{"Neiva", "America/Bogota"}
	case "NVI":
		return Location{"Navoi", "Asia/Samarkand"}
	case "NVT":
		return Location{"Navegantes", "America/Sao_Paulo"}
	case "NWI":
		return Location{"Norwich", "Europe/London"}
	case "NYA":
		return Location{"Nyagan", "Asia/Yekaterinburg"}
	case "NYI":
		return Location{"Sunyani", "Africa/Accra"}
	case "NYK":
		return Location{"Nanyuki", "Africa/Nairobi"}
	case "NYM":
		return Location{"Nadym", "Asia/Yekaterinburg"}
	case "NYO":
		return Location{"Stockholm / Nykoping", "Europe/Stockholm"}
	case "NYR":
		return Location{"Nyurba", "Asia/Yakutsk"}
	case "NYT":
		return Location{"Pyinmana", "Asia/Yangon"}
	case "NYU":
		return Location{"Nyaung U", "Asia/Yangon"}
	case "NYW":
		return Location{"Monywar", "Asia/Yangon"}
	case "OAG":
		return Location{"Orange", "Australia/Sydney"}
	case "OAJ":
		return Location{"Jacksonville", "America/New_York"}
	case "OAK":
		return Location{"Oakland", "America/Los_Angeles"}
	case "OAL":
		return Location{"Cacoal", "America/Porto_Velho"}
	case "OAX":
		return Location{"Oaxaca", "America/Mexico_City"}
	case "OBO":
		return Location{"Obihiro", "Asia/Tokyo"}
	case "OBU":
		return Location{"Kobuk", "America/Anchorage"}
	case "OCC":
		return Location{"Coca", "America/Guayaquil"}
	case "OCJ":
		return Location{"Ocho Rios", "America/Jamaica"}
	case "ODB":
		return Location{"Cordoba", "Europe/Madrid"}
	case "ODN":
		return Location{"Long Seridan", "Asia/Kuching"}
	case "ODO":
		return Location{"Bodaybo", "Asia/Irkutsk"}
	case "ODY":
		return Location{"Oudomsay", "Asia/Vientiane"}
	case "OER":
		return Location{"Ornskoldsvik", "Europe/Stockholm"}
	case "OGD":
		return Location{"Ogden", "America/Denver"}
	case "OGG":
		return Location{"Kahului", "Pacific/Honolulu"}
	case "OGL":
		return Location{"Ogle", "America/Guyana"}
	case "OGS":
		return Location{"Ogdensburg", "America/New_York"}
	case "OGU":
		return Location{"Ordu", "Europe/Istanbul"}
	case "OGX":
		return Location{"Ouargla", "Africa/Algiers"}
	case "OGZ":
		return Location{"Beslan", "Europe/Moscow"}
	case "OHD":
		return Location{"Ohrid", "Europe/Skopje"}
	case "OHE":
		return Location{"Mohe", "Asia/Shanghai"}
	case "OHH":
		return Location{"Okha", "Asia/Sakhalin"}
	case "OHO":
		return Location{"Okhotsk", "Asia/Vladivostok"}
	case "OHS":
		return Location{"Sohar", "Asia/Muscat"}
	case "OIR":
		return Location{"", "Asia/Tokyo"}
	case "OIT":
		return Location{"Oita", "Asia/Tokyo"}
	case "OKA":
		return Location{"Naha", "Asia/Tokyo"}
	case "OKC":
		return Location{"Oklahoma City", "America/Chicago"}
	case "OKD":
		return Location{"Sapporo", "Asia/Tokyo"}
	case "OKE":
		return Location{"", "Asia/Tokyo"}
	case "OKI":
		return Location{"Okinoshima", "Asia/Tokyo"}
	case "OKJ":
		return Location{"Okayama City", "Asia/Tokyo"}
	case "OKR":
		return Location{"Yorke Island", "Australia/Brisbane"}
	case "OLA":
		return Location{"Orland", "Europe/Oslo"}
	case "OLB":
		return Location{"Olbia", "Europe/Rome"}
	case "OLF":
		return Location{"Wolf Point", "America/Denver"}
	case "OLP":
		return Location{"Olympic Dam", "Australia/Adelaide"}
	case "OLZ":
		return Location{"Olyokminsk", "Asia/Yakutsk"}
	case "OMA":
		return Location{"Omaha", "America/Chicago"}
	case "OMD":
		return Location{"Oranjemund", "Africa/Johannesburg"}
	case "OME":
		return Location{"Nome", "America/Nome"}
	case "OMH":
		return Location{"Urmia", "Asia/Tehran"}
	case "OMO":
		return Location{"Mostar", "Europe/Sarajevo"}
	case "OMR":
		return Location{"Oradea", "Europe/Bucharest"}
	case "OMS":
		return Location{"Omsk", "Asia/Omsk"}
	case "OND":
		return Location{"Ondangwa", "Africa/Windhoek"}
	case "ONG":
		return Location{"", "Australia/Brisbane"}
	case "ONJ":
		return Location{"Odate", "Asia/Tokyo"}
	case "ONK":
		return Location{"Olenyok", "Asia/Yakutsk"}
	case "ONQ":
		return Location{"Zonguldak", "Europe/Istanbul"}
	case "ONS":
		return Location{"", "Australia/Perth"}
	case "ONT":
		return Location{"Ontario", "America/Los_Angeles"}
	case "OOK":
		return Location{"Toksook Bay", "America/Nome"}
	case "OOL":
		return Location{"Gold Coast", "Australia/Brisbane"}
	case "OPF":
		return Location{"Miami", "America/New_York"}
	case "OPO":
		return Location{"Porto", "Europe/Lisbon"}
	case "OPS":
		return Location{"Sinop", "America/Cuiaba"}
	case "ORB":
		return Location{"Orebro", "Europe/Stockholm"}
	case "ORD":
		return Location{"Chicago", "America/Chicago"}
	case "ORF":
		return Location{"Norfolk", "America/New_York"}
	case "ORH":
		return Location{"Worcester", "America/New_York"}
	case "ORK":
		return Location{"Cork", "Europe/Dublin"}
	case "ORN":
		return Location{"Oran", "Africa/Algiers"}
	case "ORT":
		return Location{"Northway", "America/Anchorage"}
	case "ORU":
		return Location{"Oruro", "America/La_Paz"}
	case "ORV":
		return Location{"Noorvik", "America/Anchorage"}
	case "ORX":
		return Location{"Oriximina", "America/Santarem"}
	case "ORY":
		return Location{"Paris", "Europe/Paris"}
	case "OSD":
		return Location{"Ostersund", "Europe/Stockholm"}
	case "OSI":
		return Location{"Osijek", "Europe/Zagreb"}
	case "OSL":
		return Location{"Oslo", "Europe/Oslo"}
	case "OSR":
		return Location{"Ostrava", "Europe/Prague"}
	case "OSS":
		return Location{"Osh", "Asia/Bishkek"}
	case "OST":
		return Location{"Ostend", "Europe/Brussels"}
	case "OSY":
		return Location{"Namsos", "Europe/Oslo"}
	case "OTH":
		return Location{"North Bend", "America/Los_Angeles"}
	case "OTI":
		return Location{"Gotalalamo-Morotai Island", "Asia/Jayapura"}
	case "OTP":
		return Location{"Bucharest", "Europe/Bucharest"}
	case "OTZ":
		return Location{"Kotzebue", "America/Nome"}
	case "OUA":
		return Location{"Ouagadougou", "Africa/Ouagadougou"}
	case "OUD":
		return Location{"Oujda", "Africa/Casablanca"}
	case "OUI":
		return Location{"", "Asia/Bangkok"}
	case "OUL":
		return Location{"Oulu / Oulunsalo", "Europe/Helsinki"}
	case "OUZ":
		return Location{"Zouerate", "Africa/Nouakchott"}
	case "OVB":
		return Location{"Novosibirsk", "Asia/Novosibirsk"}
	case "OVD":
		return Location{"Ranon", "Europe/Madrid"}
	case "OVS":
		return Location{"Sovetskiy", "Asia/Yekaterinburg"}
	case "OWB":
		return Location{"Owensboro", "America/Chicago"}
	case "OXB":
		return Location{"Bissau", "Africa/Bissau"}
	case "OZC":
		return Location{"Ozamiz City", "Asia/Manila"}
	case "OZG":
		return Location{"Zagora", "Africa/Casablanca"}
	case "OZZ":
		return Location{"Ouarzazate", "Africa/Casablanca"}
	case "PAB":
		return Location{"", "Asia/Kolkata"}
	case "PAC":
		return Location{"Albrook", "America/Panama"}
	case "PAD":
		return Location{"Paderborn", "Europe/Berlin"}
	case "PAE":
		return Location{"Everett", "America/Los_Angeles"}
	case "PAF":
		return Location{"", "Africa/Kampala"}
	case "PAG":
		return Location{"Pagadian City", "Asia/Manila"}
	case "PAH":
		return Location{"Paducah", "America/Chicago"}
	case "PAP":
		return Location{"Port-au-Prince", "America/Port-au-Prince"}
	case "PAS":
		return Location{"Paros Island", "Europe/Athens"}
	case "PAT":
		return Location{"Patna", "Asia/Kolkata"}
	case "PAV":
		return Location{"Paulo Afonso", "America/Bahia"}
	case "PBC":
		return Location{"Puebla", "America/Mexico_City"}
	case "PBG":
		return Location{"Plattsburgh", "America/New_York"}
	case "PBH":
		return Location{"Paro", "Asia/Thimphu"}
	case "PBI":
		return Location{"West Palm Beach", "America/New_York"}
	case "PBJ":
		return Location{"Paama Island", "Pacific/Efate"}
	case "PBM":
		return Location{"Zandery", "America/Paramaribo"}
	case "PBO":
		return Location{"Paraburdoo", "Australia/Perth"}
	case "PBR":
		return Location{"Puerto Barrios", "America/Guatemala"}
	case "PBU":
		return Location{"Putao", "Asia/Yangon"}
	case "PBZ":
		return Location{"Plettenberg Bay", "Africa/Johannesburg"}
	case "PCL":
		return Location{"Pucallpa", "America/Lima"}
	case "PCN":
		return Location{"Picton", "Pacific/Auckland"}
	case "PCP":
		return Location{"", "Africa/Sao_Tome"}
	case "PCR":
		return Location{"Puerto Carreno", "America/Bogota"}
	case "PDA":
		return Location{"Puerto Inirida", "America/Bogota"}
	case "PDG":
		return Location{"Ketaping/Padang-Sumatra Island", "Asia/Jakarta"}
	case "PDL":
		return Location{"Ponta Delgada", "Atlantic/Azores"}
	case "PDP":
		return Location{"Punta del Este", "America/Montevideo"}
	case "PDS":
		return Location{"", "America/Matamoros"}
	case "PDT":
		return Location{"Pendleton", "America/Los_Angeles"}
	case "PDV":
		return Location{"Plovdiv", "Europe/Sofia"}
	case "PDX":
		return Location{"Portland", "America/Los_Angeles"}
	case "PED":
		return Location{"Pardubice", "Europe/Prague"}
	case "PEE":
		return Location{"Perm", "Asia/Yekaterinburg"}
	case "PEG":
		return Location{"Perugia", "Europe/Rome"}
	case "PEI":
		return Location{"Pereira", "America/Bogota"}
	case "PEK":
		return Location{"Beijing", "Asia/Shanghai"}
	case "PEM":
		return Location{"Puerto Maldonado", "America/Lima"}
	case "PEN":
		return Location{"Penang", "Asia/Kuala_Lumpur"}
	case "PER":
		return Location{"Perth", "Australia/Perth"}
	case "PES":
		return Location{"Petrozavodsk", "Europe/Moscow"}
	case "PET":
		return Location{"Pelotas", "America/Sao_Paulo"}
	case "PEU":
		return Location{"Puerto Lempira", "America/Tegucigalpa"}
	case "PEW":
		return Location{"Peshawar", "Asia/Karachi"}
	case "PEZ":
		return Location{"Penza", "Europe/Moscow"}
	case "PFB":
		return Location{"Passo Fundo", "America/Sao_Paulo"}
	case "PFO":
		return Location{"Paphos", "Asia/Nicosia"}
	case "PGA":
		return Location{"Page", "America/Phoenix"}
	case "PGD":
		return Location{"Punta Gorda", "America/New_York"}
	case "PGF":
		return Location{"Perpignan/Rivesaltes", "Europe/Paris"}
	case "PGH":
		return Location{"Pantnagar", "Asia/Kolkata"}
	case "PGK":
		return Location{"Pangkal Pinang-Palaubangka Island", "Asia/Jakarta"}
	case "PGV":
		return Location{"Greenville", "America/New_York"}
	case "PGZ":
		return Location{"Ponta Grossa", "America/Sao_Paulo"}
	case "PHB":
		return Location{"Parnaiba", "America/Fortaleza"}
	case "PHC":
		return Location{"Port Harcourt", "Africa/Lagos"}
	case "PHE":
		return Location{"Port Hedland", "Australia/Perth"}
	case "PHF":
		return Location{"Newport News", "America/New_York"}
	case "PHL":
		return Location{"Philadelphia", "America/New_York"}
	case "PHO":
		return Location{"Point Hope", "America/Nome"}
	case "PHS":
		return Location{"", "Asia/Bangkok"}
	case "PHX":
		return Location{"Phoenix", "America/Phoenix"}
	case "PIA":
		return Location{"Peoria", "America/Chicago"}
	case "PIB":
		return Location{"Hattiesburg/Laurel", "America/Chicago"}
	case "PIE":
		return Location{"St Petersburg-Clearwater", "America/New_York"}
	case "PIH":
		return Location{"Pocatello", "America/Boise"}
	case "PIK":
		return Location{"Glasgow", "Europe/London"}
	case "PIN":
		return Location{"Parintins", "America/Manaus"}
	case "PIP":
		return Location{"Pilot Point", "America/Anchorage"}
	case "PIR":
		return Location{"Pierre", "America/Chicago"}
	case "PIS":
		return Location{"Poitiers/Biard", "Europe/Paris"}
	case "PIT":
		return Location{"Pittsburgh", "America/New_York"}
	case "PIU":
		return Location{"Piura", "America/Lima"}
	case "PIX":
		return Location{"Pico Island", "Atlantic/Azores"}
	case "PIZ":
		return Location{"Point Lay", "America/Nome"}
	case "PJA":
		return Location{"", "Europe/Stockholm"}
	case "PJM":
		return Location{"Puerto Jimenez", "America/Costa_Rica"}
	case "PKA":
		return Location{"Napaskiak", "America/Anchorage"}
	case "PKB":
		return Location{"Parkersburg", "America/New_York"}
	case "PKC":
		return Location{"Petropavlovsk-Kamchatsky", "Asia/Kamchatka"}
	case "PKE":
		return Location{"Parkes", "Australia/Sydney"}
	case "PKN":
		return Location{"Pangkalanbun-Borneo Island", "Asia/Pontianak"}
	case "PKP":
		return Location{"", "Pacific/Tahiti"}
	case "PKR":
		return Location{"Pokhara", "Asia/Kathmandu"}
	case "PKU":
		return Location{"Pekanbaru-Sumatra Island", "Asia/Jakarta"}
	case "PKV":
		return Location{"Pskov", "Europe/Moscow"}
	case "PKX":
		return Location{"Beijing", "Asia/Shanghai"}
	case "PKY":
		return Location{"Palangkaraya-Kalimantan Tengah", "Asia/Pontianak"}
	case "PKZ":
		return Location{"Pakse", "Asia/Vientiane"}
	case "PLM":
		return Location{"Palembang-Sumatra Island", "Asia/Jakarta"}
	case "PLN":
		return Location{"Pellston", "America/Detroit"}
	case "PLO":
		return Location{"Port Lincoln", "Australia/Adelaide"}
	case "PLQ":
		return Location{"Palanga", "Europe/Vilnius"}
	case "PLS":
		return Location{"Providenciales Island", "America/Grand_Turk"}
	case "PLW":
		return Location{"Palu-Celebes Island", "Asia/Makassar"}
	case "PLX":
		return Location{"Semey", "Asia/Almaty"}
	case "PLZ":
		return Location{"Port Elizabeth", "Africa/Johannesburg"}
	case "PMA":
		return Location{"Chake", "Africa/Dar_es_Salaam"}
	case "PMC":
		return Location{"Puerto Montt", "America/Santiago"}
	case "PMF":
		return Location{"Parma", "Europe/Rome"}
	case "PMG":
		return Location{"Ponta Pora", "America/Asuncion"}
	case "PMI":
		return Location{"Palma De Mallorca", "Europe/Madrid"}
	case "PML":
		return Location{"Cold Bay", "America/Anchorage"}
	case "PMO":
		return Location{"Palermo", "Europe/Rome"}
	case "PMR":
		return Location{"", "Pacific/Auckland"}
	case "PMV":
		return Location{"Isla Margarita", "America/Caracas"}
	case "PMW":
		return Location{"Palmas", "America/Araguaina"}
	case "PMY":
		return Location{"Puerto Madryn", "America/Argentina/Catamarca"}
	case "PNA":
		return Location{"Pamplona", "Europe/Madrid"}
	case "PNH":
		return Location{"Phnom Penh", "Asia/Phnom_Penh"}
	case "PNI":
		return Location{"Pohnpei Island", "Pacific/Pohnpei"}
	case "PNK":
		return Location{"Pontianak-Borneo Island", "Asia/Pontianak"}
	case "PNL":
		return Location{"Pantelleria", "Europe/Rome"}
	case "PNP":
		return Location{"Popondetta", "Pacific/Port_Moresby"}
	case "PNQ":
		return Location{"Pune", "Asia/Kolkata"}
	case "PNR":
		return Location{"Pointe Noire", "Africa/Brazzaville"}
	case "PNS":
		return Location{"Pensacola", "America/Chicago"}
	case "PNT":
		return Location{"Puerto Natales", "America/Punta_Arenas"}
	case "PNY":
		return Location{"", "Asia/Kolkata"}
	case "PNZ":
		return Location{"Petrolina", "America/Recife"}
	case "POA":
		return Location{"Porto Alegre", "America/Sao_Paulo"}
	case "POG":
		return Location{"Port Gentil", "Africa/Libreville"}
	case "POJ":
		return Location{"Patos De Minas", "America/Sao_Paulo"}
	case "POL":
		return Location{"Pemba / Porto Amelia", "Africa/Maputo"}
	case "POM":
		return Location{"Port Moresby", "Pacific/Port_Moresby"}
	case "POP":
		return Location{"Puerto Plata", "America/Santo_Domingo"}
	case "POR":
		return Location{"Pori", "Europe/Helsinki"}
	case "POS":
		return Location{"Port of Spain", "America/Port_of_Spain"}
	case "POZ":
		return Location{"Poznan", "Europe/Warsaw"}
	case "PPB":
		return Location{"Presidente Prudente", "America/Sao_Paulo"}
	case "PPG":
		return Location{"Pago Pago", "Pacific/Pago_Pago"}
	case "PPK":
		return Location{"Petropavlosk", "Asia/Almaty"}
	case "PPN":
		return Location{"Popayan", "America/Bogota"}
	case "PPP":
		return Location{"Proserpine", "Australia/Brisbane"}
	case "PPQ":
		return Location{"", "Pacific/Auckland"}
	case "PPS":
		return Location{"Puerto Princesa City", "Asia/Manila"}
	case "PPT":
		return Location{"Papeete", "Pacific/Tahiti"}
	case "PQC":
		return Location{"Duong Dong", "Asia/Ho_Chi_Minh"}
	case "PQI":
		return Location{"Presque Isle", "America/New_York"}
	case "PQQ":
		return Location{"Port Macquarie", "Australia/Sydney"}
	case "PRA":
		return Location{"Parana", "America/Argentina/Cordoba"}
	case "PRC":
		return Location{"Prescott", "America/Phoenix"}
	case "PRG":
		return Location{"Prague", "Europe/Prague"}
	case "PRI":
		return Location{"Praslin Island", "Indian/Mahe"}
	case "PRN":
		return Location{"Prishtina", "Europe/Belgrade"}
	case "PRS":
		return Location{"Parasi", "Pacific/Guadalcanal"}
	case "PSA":
		return Location{"Pisa", "Europe/Rome"}
	case "PSC":
		return Location{"Pasco", "America/Los_Angeles"}
	case "PSE":
		return Location{"Ponce", "America/Puerto_Rico"}
	case "PSG":
		return Location{"Petersburg", "America/Sitka"}
	case "PSM":
		return Location{"Portsmouth", "America/New_York"}
	case "PSO":
		return Location{"Pasto", "America/Bogota"}
	case "PSP":
		return Location{"Palm Springs", "America/Los_Angeles"}
	case "PSR":
		return Location{"Pescara", "Europe/Rome"}
	case "PSS":
		return Location{"Posadas", "America/Argentina/Cordoba"}
	case "PSU":
		return Location{"Putussibau-Borneo Island", "Asia/Pontianak"}
	case "PTA":
		return Location{"Port Alsworth", "America/Anchorage"}
	case "PTG":
		return Location{"Potgietersrus", "Africa/Johannesburg"}
	case "PTH":
		return Location{"Port Heiden", "America/Anchorage"}
	case "PTO":
		return Location{"Pato Branco", "America/Sao_Paulo"}
	case "PTP":
		return Location{"Pointe-a-Pitre Le Raizet", "America/Guadeloupe"}
	case "PTQ":
		return Location{"Porto De Moz", "America/Belem"}
	case "PTU":
		return Location{"Platinum", "America/Anchorage"}
	case "PTX":
		return Location{"Pitalito", "America/Bogota"}
	case "PTY":
		return Location{"Tocumen", "America/Panama"}
	case "PUB":
		return Location{"Pueblo", "America/Denver"}
	case "PUF":
		return Location{"Pau/Pyrenees (Uzein)", "Europe/Paris"}
	case "PUJ":
		return Location{"Punta Cana", "America/Santo_Domingo"}
	case "PUK":
		return Location{"Pukarua", "Pacific/Tahiti"}
	case "PUQ":
		return Location{"Punta Arenas", "America/Punta_Arenas"}
	case "PUS":
		return Location{"Busan", "Asia/Seoul"}
	case "PUU":
		return Location{"Puerto Asis", "America/Bogota"}
	case "PUW":
		return Location{"Pullman/Moscow", "America/Los_Angeles"}
	case "PUY":
		return Location{"Pula", "Europe/Zagreb"}
	case "PVA":
		return Location{"Providencia", "America/Bogota"}
	case "PVC":
		return Location{"Provincetown", "America/New_York"}
	case "PVD":
		return Location{"Providence", "America/New_York"}
	case "PVG":
		return Location{"Shanghai", "Asia/Shanghai"}
	case "PVH":
		return Location{"Porto Velho", "America/Porto_Velho"}
	case "PVK":
		return Location{"Preveza/Lefkada", "Europe/Athens"}
	case "PVR":
		return Location{"Puerto Vallarta", "America/Bahia_Banderas"}
	case "PVU":
		return Location{"Provo", "America/Denver"}
	case "PWE":
		return Location{"Pevek", "Asia/Anadyr"}
	case "PWM":
		return Location{"Portland", "America/New_York"}
	case "PWQ":
		return Location{"Pavlodar", "Asia/Almaty"}
	case "PXM":
		return Location{"Puerto Escondido", "America/Mexico_City"}
	case "PXO":
		return Location{"Vila Baleira", "Atlantic/Madeira"}
	case "PXU":
		return Location{"Pleiku", "Asia/Ho_Chi_Minh"}
	case "PYB":
		return Location{"Jeypore", "Asia/Kolkata"}
	case "PYH":
		return Location{"", "America/Bogota"}
	case "PYJ":
		return Location{"Yakutia", "Asia/Yakutsk"}
	case "PYM":
		return Location{"Plymouth", "America/New_York"}
	case "PZB":
		return Location{"Pietermaritzburg", "Africa/Johannesburg"}
	case "PZO":
		return Location{"Puerto Ordaz-Ciudad Guayana", "America/Caracas"}
	case "PZU":
		return Location{"Port Sudan", "Africa/Khartoum"}
	case "QBC":
		return Location{"Bella Coola", "America/Vancouver"}
	case "QGP":
		return Location{"Garanhuns", "America/Recife"}
	case "QIG":
		return Location{"Iguatu", "America/Fortaleza"}
	case "QOW":
		return Location{"Owerri", "Africa/Lagos"}
	case "QPA":
		return Location{"Padova", "Europe/Rome"}
	case "QRO":
		return Location{"Queretaro", "America/Mexico_City"}
	case "QSF":
		return Location{"Setif", "Africa/Algiers"}
	case "QXB":
		return Location{"Lyon", "Europe/Paris"}
	case "RAB":
		return Location{"Tokua", "Pacific/Port_Moresby"}
	case "RAE":
		return Location{"Arar", "Asia/Riyadh"}
	case "RAH":
		return Location{"Rafha", "Asia/Riyadh"}
	case "RAI":
		return Location{"Praia", "Atlantic/Cape_Verde"}
	case "RAK":
		return Location{"Marrakech", "Africa/Casablanca"}
	case "RAO":
		return Location{"Ribeirao Preto", "America/Sao_Paulo"}
	case "RAP":
		return Location{"Rapid City", "America/Denver"}
	case "RAR":
		return Location{"Avarua", "Pacific/Rarotonga"}
	case "RAS":
		return Location{"Rasht", "Asia/Tehran"}
	case "RBA":
		return Location{"Rabat", "Africa/Casablanca"}
	case "RBB":
		return Location{"Borba", "America/Manaus"}
	case "RBR":
		return Location{"Rio Branco", "America/Rio_Branco"}
	case "RBV":
		return Location{"Ramata", "Pacific/Guadalcanal"}
	case "RBY":
		return Location{"Ruby", "America/Anchorage"}
	case "RCB":
		return Location{"Richards Bay", "Africa/Johannesburg"}
	case "RCH":
		return Location{"Riohacha", "America/Bogota"}
	case "RCM":
		return Location{"", "Australia/Brisbane"}
	case "RCQ":
		return Location{"Reconquista", "America/Argentina/Cordoba"}
	case "RCU":
		return Location{"Rio Cuarto", "America/Argentina/Cordoba"}
	case "RDD":
		return Location{"Redding", "America/Los_Angeles"}
	case "RDM":
		return Location{"Redmond", "America/Los_Angeles"}
	case "RDO":
		return Location{"Radom", "Europe/Warsaw"}
	case "RDU":
		return Location{"Raleigh/Durham", "America/New_York"}
	case "RDZ":
		return Location{"Rodez/Marcillac", "Europe/Paris"}
	case "REA":
		return Location{"", "Pacific/Tahiti"}
	case "REC":
		return Location{"Recife", "America/Recife"}
	case "REG":
		return Location{"Reggio Calabria", "Europe/Rome"}
	case "REL":
		return Location{"Rawson", "America/Argentina/Catamarca"}
	case "REN":
		return Location{"Orenburg", "Asia/Yekaterinburg"}
	case "RER":
		return Location{"Retalhuleu", "America/Guatemala"}
	case "RES":
		return Location{"Resistencia", "America/Argentina/Cordoba"}
	case "RET":
		return Location{"", "Europe/Oslo"}
	case "REU":
		return Location{"Reus", "Europe/Madrid"}
	case "REX":
		return Location{"Reynosa", "America/Matamoros"}
	case "RFD":
		return Location{"Chicago/Rockford", "America/Chicago"}
	case "RFP":
		return Location{"Uturoa", "Pacific/Tahiti"}
	case "RGA":
		return Location{"Rio Grande", "America/Argentina/Ushuaia"}
	case "RGI":
		return Location{"", "Pacific/Tahiti"}
	case "RGK":
		return Location{"Gorno-Altaysk", "Asia/Barnaul"}
	case "RGL":
		return Location{"Rio Gallegos", "America/Argentina/Rio_Gallegos"}
	case "RGN":
		return Location{"Yangon", "Asia/Yangon"}
	case "RGS":
		return Location{"Burgos", "Europe/Madrid"}
	case "RHD":
		return Location{"Rio Hondo", "America/Argentina/Cordoba"}
	case "RHI":
		return Location{"Rhinelander", "America/Chicago"}
	case "RHO":
		return Location{"Rodes Island", "Europe/Athens"}
	case "RIA":
		return Location{"Santa Maria", "America/Sao_Paulo"}
	case "RIC":
		return Location{"Richmond", "America/New_York"}
	case "RIG":
		return Location{"Rio Grande", "America/Sao_Paulo"}
	case "RIH":
		return Location{"Rio Hato", "America/Panama"}
	case "RIS":
		return Location{"Rishiri", "Asia/Tokyo"}
	case "RIW":
		return Location{"Riverton", "America/Denver"}
	case "RIX":
		return Location{"Riga", "Europe/Riga"}
	case "RJA":
		return Location{"Rajahmundry", "Asia/Kolkata"}
	case "RJB":
		return Location{"Rajbiraj", "Asia/Kathmandu"}
	case "RJH":
		return Location{"Rajshahi", "Asia/Dhaka"}
	case "RJK":
		return Location{"Rijeka", "Europe/Zagreb"}
	case "RJL":
		return Location{"Logrono", "Europe/Madrid"}
	case "RKA":
		return Location{"", "Pacific/Tahiti"}
	case "RKD":
		return Location{"Rockland", "America/New_York"}
	case "RKS":
		return Location{"Rock Springs", "America/Denver"}
	case "RKT":
		return Location{"Ras Al Khaimah", "Asia/Dubai"}
	case "RKV":
		return Location{"Reykjavik", "Atlantic/Reykjavik"}
	case "RLG":
		return Location{"Rostock", "Europe/Berlin"}
	case "RLO":
		return Location{"Merlo", "America/Argentina/San_Luis"}
	case "RMA":
		return Location{"Roma", "Australia/Brisbane"}
	case "RMF":
		return Location{"Marsa Alam", "Africa/Cairo"}
	case "RMI":
		return Location{"Rimini", "Europe/Rome"}
	case "RMQ":
		return Location{"Taichung City", "Asia/Taipei"}
	case "RMU":
		return Location{"Corvera", "Europe/Madrid"}
	case "RNA":
		return Location{"Arona", "Pacific/Guadalcanal"}
	case "RNB":
		return Location{"", "Europe/Stockholm"}
	case "RNJ":
		return Location{"", "Asia/Tokyo"}
	case "RNL":
		return Location{"Rennell Island", "Pacific/Guadalcanal"}
	case "RNN":
		return Location{"Ronne", "Europe/Copenhagen"}
	case "RNO":
		return Location{"Reno", "America/Los_Angeles"}
	case "RNS":
		return Location{"Rennes/Saint-Jacques", "Europe/Paris"}
	case "ROA":
		return Location{"Roanoke", "America/New_York"}
	case "ROB":
		return Location{"Monrovia", "Africa/Monrovia"}
	case "ROC":
		return Location{"Rochester", "America/New_York"}
	case "ROI":
		return Location{"", "Asia/Bangkok"}
	case "ROK":
		return Location{"Rockhampton", "Australia/Brisbane"}
	case "RON":
		return Location{"Paipa", "America/Bogota"}
	case "ROO":
		return Location{"Rondonopolis", "America/Cuiaba"}
	case "ROR":
		return Location{"Babelthuap Island", "Pacific/Palau"}
	case "ROS":
		return Location{"Rosario", "America/Argentina/Cordoba"}
	case "ROT":
		return Location{"Rotorua", "Pacific/Auckland"}
	case "ROV":
		return Location{"Rostov-on-Don", "Europe/Moscow"}
	case "ROW":
		return Location{"Roswell", "America/Denver"}
	case "RPR":
		return Location{"Raipur", "Asia/Kolkata"}
	case "RRG":
		return Location{"Port Mathurin", "Indian/Mauritius"}
	case "RRK":
		return Location{"", "Asia/Kolkata"}
	case "RRR":
		return Location{"", "Pacific/Tahiti"}
	case "RRS":
		return Location{"Roros", "Europe/Oslo"}
	case "RSA":
		return Location{"Santa Rosa", "America/Argentina/Salta"}
	case "RSD":
		return Location{"Rock Sound", "America/Nassau"}
	case "RSH":
		return Location{"Russian Mission", "America/Anchorage"}
	case "RST":
		return Location{"Rochester", "America/Chicago"}
	case "RSU":
		return Location{"Yeosu", "Asia/Seoul"}
	case "RSW":
		return Location{"Fort Myers", "America/New_York"}
	case "RTA":
		return Location{"Rotuma", "Pacific/Fiji"}
	case "RTB":
		return Location{"Roatan Island", "America/Tegucigalpa"}
	case "RTG":
		return Location{"Satar Tacik-Flores Island", "Asia/Makassar"}
	case "RTM":
		return Location{"Rotterdam", "Europe/Amsterdam"}
	case "RUH":
		return Location{"Riyadh", "Asia/Riyadh"}
	case "RUN":
		return Location{"St Denis", "Indian/Reunion"}
	case "RUR":
		return Location{"", "Pacific/Tahiti"}
	case "RUT":
		return Location{"Rutland", "America/New_York"}
	case "RVD":
		return Location{"Rio Verde", "America/Sao_Paulo"}
	case "RVE":
		return Location{"Saravena", "America/Bogota"}
	case "RVK":
		return Location{"Rorvik", "Europe/Oslo"}
	case "RVN":
		return Location{"Rovaniemi", "Europe/Helsinki"}
	case "RVV":
		return Location{"", "Pacific/Tahiti"}
	case "RXS":
		return Location{"Roxas City", "Asia/Manila"}
	case "RZE":
		return Location{"Rzeszow", "Europe/Warsaw"}
	case "RZR":
		return Location{"", "Asia/Tehran"}
	case "SAB":
		return Location{"Saba", "America/Kralendijk"}
	case "SAF":
		return Location{"Santa Fe", "America/Denver"}
	case "SAG":
		return Location{"Kakadi", "Asia/Kolkata"}
	case "SAL":
		return Location{"Santa Clara", "America/El_Salvador"}
	case "SAN":
		return Location{"San Diego", "America/Los_Angeles"}
	case "SAP":
		return Location{"La Mesa", "America/Tegucigalpa"}
	case "SAQ":
		return Location{"Andros Island", "America/Nassau"}
	case "SAT":
		return Location{"San Antonio", "America/Chicago"}
	case "SAV":
		return Location{"Savannah", "America/New_York"}
	case "SAW":
		return Location{"Istanbul", "Europe/Istanbul"}
	case "SBA":
		return Location{"Santa Barbara", "America/Los_Angeles"}
	case "SBD":
		return Location{"San Bernardino", "America/Los_Angeles"}
	case "SBH":
		return Location{"Gustavia", "America/St_Barthelemy"}
	case "SBN":
		return Location{"South Bend", "America/Indiana/Indianapolis"}
	case "SBP":
		return Location{"San Luis Obispo", "America/Los_Angeles"}
	case "SBR":
		return Location{"Saibai Island", "Australia/Brisbane"}
	case "SBW":
		return Location{"Sibu", "Asia/Kuching"}
	case "SBY":
		return Location{"Salisbury", "America/New_York"}
	case "SBZ":
		return Location{"Sibiu", "Europe/Bucharest"}
	case "SCC":
		return Location{"Deadhorse", "America/Anchorage"}
	case "SCE":
		return Location{"State College", "America/New_York"}
	case "SCF":
		return Location{"Scottsdale", "America/Phoenix"}
	case "SCK":
		return Location{"Stockton", "America/Los_Angeles"}
	case "SCL":
		return Location{"Santiago", "America/Santiago"}
	case "SCM":
		return Location{"Scammon Bay", "America/Nome"}
	case "SCN":
		return Location{"Saarbrucken", "Europe/Berlin"}
	case "SCO":
		return Location{"Aktau", "Asia/Aqtau"}
	case "SCQ":
		return Location{"Santiago de Compostela", "Europe/Madrid"}
	case "SCU":
		return Location{"Santiago", "America/Havana"}
	case "SCV":
		return Location{"Suceava", "Europe/Bucharest"}
	case "SCW":
		return Location{"Syktyvkar", "Europe/Moscow"}
	case "SCY":
		return Location{"San Cristobal", "Pacific/Galapagos"}
	case "SCZ":
		return Location{"Santa Cruz/Graciosa Bay/Luova", "Pacific/Guadalcanal"}
	case "SDD":
		return Location{"Lubango", "Africa/Luanda"}
	case "SDE":
		return Location{"Santiago del Estero", "America/Argentina/Cordoba"}
	case "SDF":
		return Location{"Louisville", "America/Kentucky/Louisville"}
	case "SDG":
		return Location{"", "Asia/Tehran"}
	case "SDJ":
		return Location{"Sendai", "Asia/Tokyo"}
	case "SDK":
		return Location{"Sandakan", "Asia/Kuching"}
	case "SDL":
		return Location{"Sundsvall/ Harnosand", "Europe/Stockholm"}
	case "SDN":
		return Location{"Sandane", "Europe/Oslo"}
	case "SDP":
		return Location{"Sand Point", "America/Anchorage"}
	case "SDQ":
		return Location{"Santo Domingo", "America/Santo_Domingo"}
	case "SDR":
		return Location{"Santander", "Europe/Madrid"}
	case "SDU":
		return Location{"Rio De Janeiro", "America/Sao_Paulo"}
	case "SDY":
		return Location{"Sidney", "America/Denver"}
	case "SEA":
		return Location{"Seattle", "America/Los_Angeles"}
	case "SEB":
		return Location{"Sabha", "Africa/Tripoli"}
	case "SEN":
		return Location{"Southend", "Europe/London"}
	case "SEU":
		return Location{"Seronera", "Africa/Dar_es_Salaam"}
	case "SEZ":
		return Location{"Mahe Island", "Indian/Mahe"}
	case "SFA":
		return Location{"Sfax", "Africa/Tunis"}
	case "SFB":
		return Location{"Orlando", "America/New_York"}
	case "SFD":
		return Location{"Inglaterra", "America/Caracas"}
	case "SFG":
		return Location{"Grand Case", "America/Lower_Princes"}
	case "SFJ":
		return Location{"Kangerlussuaq", "America/Nuuk"}
	case "SFL":
		return Location{"Sao Filipe", "Atlantic/Cape_Verde"}
	case "SFN":
		return Location{"Santa Fe", "America/Argentina/Cordoba"}
	case "SFO":
		return Location{"San Francisco", "America/Los_Angeles"}
	case "SFT":
		return Location{"Skelleftea", "Europe/Stockholm"}
	case "SGC":
		return Location{"Surgut", "Asia/Yekaterinburg"}
	case "SGD":
		return Location{"Sonderborg", "Europe/Copenhagen"}
	case "SGF":
		return Location{"Springfield", "America/Chicago"}
	case "SGN":
		return Location{"Ho Chi Minh City", "Asia/Ho_Chi_Minh"}
	case "SGO":
		return Location{"", "Australia/Brisbane"}
	case "SGU":
		return Location{"St George", "America/Denver"}
	case "SGX":
		return Location{"Songea", "Africa/Dar_es_Salaam"}
	case "SGY":
		return Location{"Skagway", "America/Juneau"}
	case "SHA":
		return Location{"Shanghai", "Asia/Shanghai"}
	case "SHB":
		return Location{"Nakashibetsu", "Asia/Tokyo"}
	case "SHC":
		return Location{"Shire", "Africa/Addis_Ababa"}
	case "SHD":
		return Location{"Staunton/Waynesboro/Harrisonburg", "America/New_York"}
	case "SHE":
		return Location{"Shenyang", "Asia/Shanghai"}
	case "SHH":
		return Location{"Shishmaref", "America/Nome"}
	case "SHI":
		return Location{"", "Asia/Tokyo"}
	case "SHJ":
		return Location{"Sharjah", "Asia/Dubai"}
	case "SHL":
		return Location{"Shillong", "Asia/Kolkata"}
	case "SHM":
		return Location{"Shirahama", "Asia/Tokyo"}
	case "SHO":
		return Location{"", "Asia/Seoul"}
	case "SHR":
		return Location{"Sheridan", "America/Denver"}
	case "SHS":
		return Location{"Shashi", "Asia/Shanghai"}
	case "SHV":
		return Location{"Shreveport", "America/Chicago"}
	case "SHW":
		return Location{"", "Asia/Riyadh"}
	case "SHX":
		return Location{"Shageluk", "America/Anchorage"}
	case "SID":
		return Location{"Espargos", "Atlantic/Cape_Verde"}
	case "SIF":
		return Location{"Simara", "Asia/Kathmandu"}
	case "SIG":
		return Location{"San Juan", "America/Puerto_Rico"}
	case "SIN":
		return Location{"Singapore", "Asia/Singapore"}
	case "SIR":
		return Location{"Sion", "Europe/Zurich"}
	case "SIS":
		return Location{"Sishen", "Africa/Johannesburg"}
	case "SIT":
		return Location{"Sitka", "America/Sitka"}
	case "SJC":
		return Location{"San Jose", "America/Los_Angeles"}
	case "SJD":
		return Location{"San Jose del Cabo", "America/Mazatlan"}
	case "SJE":
		return Location{"San Jose Del Guaviare", "America/Bogota"}
	case "SJI":
		return Location{"San Jose", "Asia/Manila"}
	case "SJJ":
		return Location{"Sarajevo", "Europe/Sarajevo"}
	case "SJK":
		return Location{"Sao Jose Dos Campos", "America/Sao_Paulo"}
	case "SJL":
		return Location{"Sao Gabriel Da Cachoeira", "America/Manaus"}
	case "SJO":
		return Location{"San Jose", "America/Costa_Rica"}
	case "SJP":
		return Location{"Sao Jose Do Rio Preto", "America/Sao_Paulo"}
	case "SJT":
		return Location{"San Angelo", "America/Chicago"}
	case "SJU":
		return Location{"San Juan", "America/Puerto_Rico"}
	case "SJW":
		return Location{"Shijiazhuang", "Asia/Shanghai"}
	case "SJZ":
		return Location{"Velas", "Atlantic/Azores"}
	case "SKB":
		return Location{"Basseterre", "America/St_Kitts"}
	case "SKD":
		return Location{"Samarkand", "Asia/Samarkand"}
	case "SKG":
		return Location{"Thessaloniki", "Europe/Athens"}
	case "SKH":
		return Location{"Surkhet", "Asia/Kathmandu"}
	case "SKK":
		return Location{"Shaktoolik", "America/Anchorage"}
	case "SKN":
		return Location{"Hadsel", "Europe/Oslo"}
	case "SKO":
		return Location{"Sokoto", "Africa/Lagos"}
	case "SKP":
		return Location{"Skopje", "Europe/Skopje"}
	case "SKT":
		return Location{"Sialkot", "Asia/Karachi"}
	case "SKU":
		return Location{"Skiros Island", "Europe/Athens"}
	case "SKX":
		return Location{"Saransk", "Europe/Moscow"}
	case "SKZ":
		return Location{"Mirpur Khas", "Asia/Karachi"}
	case "SLA":
		return Location{"Salta", "America/Argentina/Salta"}
	case "SLC":
		return Location{"Salt Lake City", "America/Denver"}
	case "SLE":
		return Location{"Salem", "America/Los_Angeles"}
	case "SLH":
		return Location{"Sola", "Pacific/Efate"}
	case "SLI":
		return Location{"Solwesi", "Africa/Lusaka"}
	case "SLK":
		return Location{"Saranac Lake", "America/New_York"}
	case "SLL":
		return Location{"Salalah", "Asia/Muscat"}
	case "SLM":
		return Location{"Salamanca", "Europe/Madrid"}
	case "SLN":
		return Location{"Salina", "America/Chicago"}
	case "SLP":
		return Location{"San Luis Potosi", "America/Mexico_City"}
	case "SLU":
		return Location{"Castries", "America/St_Lucia"}
	case "SLV":
		return Location{"", "Asia/Kolkata"}
	case "SLX":
		return Location{"Salt Cay", "America/Grand_Turk"}
	case "SLY":
		return Location{"Salekhard", "Asia/Yekaterinburg"}
	case "SLZ":
		return Location{"Sao Luis", "America/Fortaleza"}
	case "SMA":
		return Location{"Vila do Porto", "Atlantic/Azores"}
	case "SMF":
		return Location{"Sacramento", "America/Los_Angeles"}
	case "SMI":
		return Location{"Samos Island", "Europe/Athens"}
	case "SMK":
		return Location{"St Michael", "America/Nome"}
	case "SML":
		return Location{"Stella Maris", "America/Nassau"}
	case "SMQ":
		return Location{"Sampit-Borneo Island", "Asia/Pontianak"}
	case "SMR":
		return Location{"Santa Marta", "America/Bogota"}
	case "SMS":
		return Location{"", "Indian/Antananarivo"}
	case "SMX":
		return Location{"Santa Maria", "America/Los_Angeles"}
	case "SNA":
		return Location{"Santa Ana", "America/Los_Angeles"}
	case "SNE":
		return Location{"Preguica", "Atlantic/Cape_Verde"}
	case "SNN":
		return Location{"Shannon", "Europe/Dublin"}
	case "SNO":
		return Location{"", "Asia/Bangkok"}
	case "SNP":
		return Location{"St Paul Island", "America/Nome"}
	case "SNR":
		return Location{"Saint-Nazaire/Montoir", "Europe/Paris"}
	case "SNU":
		return Location{"Santa Clara", "America/Havana"}
	case "SNW":
		return Location{"Thandwe", "Asia/Yangon"}
	case "SOC":
		return Location{"Sukarata(Solo)-Java Island", "Asia/Jakarta"}
	case "SOF":
		return Location{"Sofia", "Europe/Sofia"}
	case "SOG":
		return Location{"Sogndal", "Europe/Oslo"}
	case "SOJ":
		return Location{"Sorkjosen", "Europe/Oslo"}
	case "SON":
		return Location{"Luganville", "Pacific/Efate"}
	case "SOQ":
		return Location{"Sorong-Papua Island", "Asia/Jayapura"}
	case "SOU":
		return Location{"Southampton", "Europe/London"}
	case "SOV":
		return Location{"Seldovia", "America/Anchorage"}
	case "SOW":
		return Location{"Show Low", "America/Phoenix"}
	case "SPC":
		return Location{"Sta Cruz de la Palma", "Atlantic/Canary"}
	case "SPD":
		return Location{"Saidpur", "Asia/Dhaka"}
	case "SPI":
		return Location{"Springfield", "America/Chicago"}
	case "SPN":
		return Location{"Saipan Island", "Pacific/Saipan"}
	case "SPP":
		return Location{"Menongue", "Africa/Luanda"}
	case "SPS":
		return Location{"Wichita Falls", "America/Chicago"}
	case "SPU":
		return Location{"Split", "Europe/Zagreb"}
	case "SPX":
		return Location{"Giza", "Africa/Cairo"}
	case "SPY":
		return Location{"", "Africa/Abidjan"}
	case "SQD":
		return Location{"San Francisco de Yeso", "America/Lima"}
	case "SQG":
		return Location{"Sintang-Borneo Island", "Asia/Pontianak"}
	case "SQJ":
		return Location{"Sanming", "Asia/Shanghai"}
	case "SQL":
		return Location{"San Carlos", "America/Los_Angeles"}
	case "SRA":
		return Location{"Santa Rosa", "America/Sao_Paulo"}
	case "SRE":
		return Location{"Sucre", "America/La_Paz"}
	case "SRG":
		return Location{"Semarang-Java Island", "Asia/Jakarta"}
	case "SRP":
		return Location{"Svea", "Arctic/Longyearbyen"}
	case "SRQ":
		return Location{"Sarasota/Bradenton", "America/New_York"}
	case "SRY":
		return Location{"Sari", "Asia/Tehran"}
	case "SSA":
		return Location{"Salvador", "America/Bahia"}
	case "SSG":
		return Location{"Malabo", "Africa/Malabo"}
	case "SSH":
		return Location{"Sharm el-Sheikh", "Africa/Cairo"}
	case "SSJ":
		return Location{"Alstahaug", "Europe/Oslo"}
	case "SSR":
		return Location{"Pentecost Island", "Pacific/Efate"}
	case "STB":
		return Location{"", "America/Caracas"}
	case "STC":
		return Location{"St Cloud", "America/Chicago"}
	case "STD":
		return Location{"Santo Domingo", "America/Caracas"}
	case "STI":
		return Location{"Santiago", "America/Santo_Domingo"}
	case "STL":
		return Location{"St Louis", "America/Chicago"}
	case "STM":
		return Location{"Santarem", "America/Santarem"}
	case "STN":
		return Location{"London", "Europe/London"}
	case "STR":
		return Location{"Stuttgart", "Europe/Berlin"}
	case "STS":
		return Location{"Santa Rosa", "America/Los_Angeles"}
	case "STT":
		return Location{"Charlotte Amalie", "America/St_Thomas"}
	case "STV":
		return Location{"", "Asia/Kolkata"}
	case "STW":
		return Location{"Stavropol", "Europe/Moscow"}
	case "STX":
		return Location{"Christiansted", "America/St_Thomas"}
	case "SUB":
		return Location{"Surabaya", "Asia/Jakarta"}
	case "SUF":
		return Location{"Lamezia Terme", "Europe/Rome"}
	case "SUG":
		return Location{"Surigao City", "Asia/Manila"}
	case "SUJ":
		return Location{"Satu Mare", "Europe/Bucharest"}
	case "SUN":
		return Location{"Hailey", "America/Boise"}
	case "SUV":
		return Location{"Nausori", "Pacific/Fiji"}
	case "SUX":
		return Location{"Sioux City", "America/Chicago"}
	case "SUY":
		return Location{"Suntar", "Asia/Yakutsk"}
	case "SVA":
		return Location{"Savoonga", "America/Nome"}
	case "SVB":
		return Location{"", "Indian/Antananarivo"}
	case "SVC":
		return Location{"Silver City", "America/Denver"}
	case "SVD":
		return Location{"Kingstown", "America/St_Vincent"}
	case "SVG":
		return Location{"Stavanger", "Europe/Oslo"}
	case "SVI":
		return Location{"San Vicente Del Caguan", "America/Bogota"}
	case "SVJ":
		return Location{"Svolvaer", "Europe/Oslo"}
	case "SVL":
		return Location{"Savonlinna", "Europe/Helsinki"}
	case "SVO":
		return Location{"Moscow", "Europe/Moscow"}
	case "SVQ":
		return Location{"Sevilla", "Europe/Madrid"}
	case "SVU":
		return Location{"Savusavu", "Pacific/Fiji"}
	case "SVX":
		return Location{"Yekaterinburg", "Asia/Yekaterinburg"}
	case "SVZ":
		return Location{"", "America/Bogota"}
	case "SWA":
		return Location{"Shantou", "Asia/Shanghai"}
	case "SWF":
		return Location{"Newburgh", "America/New_York"}
	case "SWJ":
		return Location{"Malekula Island", "Pacific/Efate"}
	case "SWO":
		return Location{"Stillwater", "America/Chicago"}
	case "SWQ":
		return Location{"Sumbawa Island", "Asia/Makassar"}
	case "SWT":
		return Location{"Strezhevoy", "Asia/Tomsk"}
	case "SWV":
		return Location{"Evensk", "Asia/Magadan"}
	case "SXB":
		return Location{"Strasbourg", "Europe/Paris"}
	case "SXK":
		return Location{"Saumlaki-Yamdena Island", "Asia/Jayapura"}
	case "SXM":
		return Location{"Saint Martin", "America/Lower_Princes"}
	case "SXR":
		return Location{"Srinagar", "Asia/Kolkata"}
	case "SXV":
		return Location{"", "Asia/Kolkata"}
	case "SXZ":
		return Location{"Siirt", "Europe/Istanbul"}
	case "SYD":
		return Location{"Sydney", "Australia/Sydney"}
	case "SYJ":
		return Location{"", "Asia/Tehran"}
	case "SYM":
		return Location{"Simao", "Asia/Shanghai"}
	case "SYO":
		return Location{"Shonai", "Asia/Tokyo"}
	case "SYR":
		return Location{"Syracuse", "America/New_York"}
	case "SYS":
		return Location{"Saskylakh", "Asia/Yakutsk"}
	case "SYU":
		return Location{"Sue Islet", "Australia/Brisbane"}
	case "SYX":
		return Location{"Sanya", "Asia/Shanghai"}
	case "SYY":
		return Location{"Stornoway", "Europe/London"}
	case "SYZ":
		return Location{"Shiraz", "Asia/Tehran"}
	case "SZA":
		return Location{"Soyo", "Africa/Luanda"}
	case "SZB":
		return Location{"Subang", "Asia/Kuala_Lumpur"}
	case "SZE":
		return Location{"Semera", "Africa/Addis_Ababa"}
	case "SZF":
		return Location{"Samsun", "Europe/Istanbul"}
	case "SZG":
		return Location{"Salzburg", "Europe/Berlin"}
	case "SZK":
		return Location{"Skukuza", "Africa/Johannesburg"}
	case "SZX":
		return Location{"Shenzhen", "Asia/Shanghai"}
	case "SZY":
		return Location{"Szymany", "Europe/Warsaw"}
	case "SZZ":
		return Location{"Goleniow", "Europe/Warsaw"}
	case "TAB":
		return Location{"Scarborough", "America/Port_of_Spain"}
	case "TAC":
		return Location{"Tacloban City", "Asia/Manila"}
	case "TAE":
		return Location{"Daegu", "Asia/Seoul"}
	case "TAG":
		return Location{"Tagbilaran City", "Asia/Manila"}
	case "TAH":
		return Location{"", "Pacific/Efate"}
	case "TAK":
		return Location{"Takamatsu", "Asia/Tokyo"}
	case "TAL":
		return Location{"Tanana", "America/Anchorage"}
	case "TAM":
		return Location{"Tampico", "America/Monterrey"}
	case "TAO":
		return Location{"Qingdao", "Asia/Shanghai"}
	case "TAP":
		return Location{"Tapachula", "America/Mexico_City"}
	case "TAS":
		return Location{"Tashkent", "Asia/Tashkent"}
	case "TAT":
		return Location{"Poprad", "Europe/Bratislava"}
	case "TAY":
		return Location{"Tartu", "Europe/Tallinn"}
	case "TAZ":
		return Location{"Dashoguz", "Asia/Ashgabat"}
	case "TBB":
		return Location{"Tuy Hoa", "Asia/Ho_Chi_Minh"}
	case "TBG":
		return Location{"Tabubil", "Pacific/Port_Moresby"}
	case "TBH":
		return Location{"Romblon", "Asia/Manila"}
	case "TBN":
		return Location{"Fort Leonard Wood", "America/Chicago"}
	case "TBO":
		return Location{"Tabora", "Africa/Dar_es_Salaam"}
	case "TBP":
		return Location{"Tumbes", "America/Lima"}
	case "TBS":
		return Location{"Tbilisi", "Asia/Tbilisi"}
	case "TBT":
		return Location{"Tabatinga", "America/Bogota"}
	case "TBU":
		return Location{"Nuku'alofa", "Pacific/Tongatapu"}
	case "TBW":
		return Location{"Tambov", "Europe/Moscow"}
	case "TBZ":
		return Location{"Tabriz", "Asia/Tehran"}
	case "TCA":
		return Location{"Tennant Creek", "Australia/Darwin"}
	case "TCG":
		return Location{"Tacheng", "Asia/Shanghai"}
	case "TCO":
		return Location{"Tumaco", "America/Bogota"}
	case "TCQ":
		return Location{"Tacna", "America/Lima"}
	case "TCR":
		return Location{"Thoothukkudi", "Asia/Kolkata"}
	case "TCZ":
		return Location{"Tengchong", "Asia/Shanghai"}
	case "TDD":
		return Location{"Trinidad", "America/La_Paz"}
	case "TDK":
		return Location{"Taldy Kurgan", "Asia/Almaty"}
	case "TDX":
		return Location{"", "Asia/Bangkok"}
	case "TEC":
		return Location{"Telemaco Borba", "America/Sao_Paulo"}
	case "TEE":
		return Location{"Tebessi", "Africa/Algiers"}
	case "TEI":
		return Location{"Tezu", "Asia/Kolkata"}
	case "TEN":
		return Location{"", "Asia/Shanghai"}
	case "TEQ":
		return Location{"Corlu", "Europe/Istanbul"}
	case "TER":
		return Location{"Lajes", "Atlantic/Azores"}
	case "TET":
		return Location{"Tete", "Africa/Maputo"}
	case "TEX":
		return Location{"Telluride", "America/Denver"}
	case "TEZ":
		return Location{"", "Asia/Kolkata"}
	case "TFF":
		return Location{"Tefe", "America/Manaus"}
	case "TFL":
		return Location{"Teofilo Otoni", "America/Sao_Paulo"}
	case "TFN":
		return Location{"Tenerife Island", "Atlantic/Canary"}
	case "TFS":
		return Location{"Tenerife Island", "Atlantic/Canary"}
	case "TFU":
		return Location{"Chengdu", "Asia/Shanghai"}
	case "TGC":
		return Location{"Tanjung Manis", "Asia/Kuching"}
	case "TGD":
		return Location{"Podgorica", "Europe/Podgorica"}
	case "TGG":
		return Location{"Kuala Terengganu", "Asia/Kuala_Lumpur"}
	case "TGI":
		return Location{"Tingo Maria", "America/Lima"}
	case "TGM":
		return Location{"Targu Mures", "Europe/Bucharest"}
	case "TGO":
		return Location{"Tongliao", "Asia/Shanghai"}
	case "TGP":
		return Location{"Bor", "Asia/Krasnoyarsk"}
	case "TGR":
		return Location{"Touggourt", "Africa/Algiers"}
	case "TGT":
		return Location{"Tanga", "Africa/Dar_es_Salaam"}
	case "TGU":
		return Location{"Tegucigalpa", "America/Tegucigalpa"}
	case "TGZ":
		return Location{"Tuxtla Gutierrez", "America/Mexico_City"}
	case "THD":
		return Location{"Thanh Hóa", "Asia/Bangkok"}
	case "THE":
		return Location{"Teresina", "America/Fortaleza"}
	case "THL":
		return Location{"Tachileik", "Asia/Yangon"}
	case "THN":
		return Location{"Trollhattan", "Europe/Stockholm"}
	case "THQ":
		return Location{"Tianshui", "Asia/Shanghai"}
	case "THR":
		return Location{"Tehran", "Asia/Tehran"}
	case "THS":
		return Location{"", "Asia/Bangkok"}
	case "THU":
		return Location{"Thule", "America/Thule"}
	case "THX":
		return Location{"Turukhansk", "Asia/Krasnoyarsk"}
	case "TIA":
		return Location{"Tirana", "Europe/Tirane"}
	case "TID":
		return Location{"Tiaret", "Africa/Algiers"}
	case "TIF":
		return Location{"", "Asia/Riyadh"}
	case "TIH":
		return Location{"", "Pacific/Tahiti"}
	case "TIJ":
		return Location{"Tijuana", "America/Los_Angeles"}
	case "TIM":
		return Location{"Timika-Papua Island", "Asia/Jayapura"}
	case "TIN":
		return Location{"Tindouf", "Africa/Algiers"}
	case "TIP":
		return Location{"Tripoli", "Africa/Tripoli"}
	case "TIR":
		return Location{"Tirupati", "Asia/Kolkata"}
	case "TIU":
		return Location{"", "Pacific/Auckland"}
	case "TIV":
		return Location{"Tivat", "Europe/Podgorica"}
	case "TIZ":
		return Location{"Tari", "Pacific/Port_Moresby"}
	case "TJA":
		return Location{"Tarija", "America/La_Paz"}
	case "TJH":
		return Location{"Tajima", "Asia/Tokyo"}
	case "TJK":
		return Location{"Tokat", "Europe/Istanbul"}
	case "TJL":
		return Location{"Tres Lagoas", "America/Campo_Grande"}
	case "TJM":
		return Location{"Tyumen", "Asia/Yekaterinburg"}
	case "TJN":
		return Location{"Takume", "Pacific/Tahiti"}
	case "TJQ":
		return Location{"Tanjung Pandan-Belitung Island", "Asia/Jakarta"}
	case "TJS":
		return Location{"Tanjung Selor-Borneo Island", "Asia/Makassar"}
	case "TJU":
		return Location{"Kulyab", "Asia/Dushanbe"}
	case "TKD":
		return Location{"Sekondi-Takoradi", "Africa/Accra"}
	case "TKF":
		return Location{"Truckee", "America/Los_Angeles"}
	case "TKG":
		return Location{"Bandar Lampung-Sumatra Island", "Asia/Jakarta"}
	case "TKJ":
		return Location{"Tok", "America/Anchorage"}
	case "TKK":
		return Location{"Weno Island", "Pacific/Chuuk"}
	case "TKN":
		return Location{"Tokunoshima", "Asia/Tokyo"}
	case "TKP":
		return Location{"", "Pacific/Tahiti"}
	case "TKQ":
		return Location{"Kigoma", "Africa/Dar_es_Salaam"}
	case "TKS":
		return Location{"Tokushima", "Asia/Tokyo"}
	case "TKU":
		return Location{"Turku", "Europe/Helsinki"}
	case "TKV":
		return Location{"Tatakoto", "Pacific/Tahiti"}
	case "TKX":
		return Location{"", "Pacific/Tahiti"}
	case "TLA":
		return Location{"Teller", "America/Nome"}
	case "TLC":
		return Location{"Toluca", "America/Mexico_City"}
	case "TLE":
		return Location{"", "Indian/Antananarivo"}
	case "TLH":
		return Location{"Tallahassee", "America/New_York"}
	case "TLI":
		return Location{"Toli Toli-Celebes Island", "Asia/Makassar"}
	case "TLK":
		return Location{"", "Asia/Yakutsk"}
	case "TLL":
		return Location{"Tallinn", "Europe/Tallinn"}
	case "TLM":
		return Location{"Tlemcen", "Africa/Algiers"}
	case "TLN":
		return Location{"Toulon/Hyeres/Le Palyvestre", "Europe/Paris"}
	case "TLS":
		return Location{"Toulouse/Blagnac", "Europe/Paris"}
	case "TLU":
		return Location{"Tolu", "America/Bogota"}
	case "TLV":
		return Location{"Tel Aviv", "Asia/Jerusalem"}
	case "TLY":
		return Location{"Plastun", "Asia/Vladivostok"}
	case "TMC":
		return Location{"Waikabubak-Sumba Island", "Asia/Makassar"}
	case "TME":
		return Location{"Tame", "America/Bogota"}
	case "TMI":
		return Location{"Tumling Tar", "Asia/Kathmandu"}
	case "TMJ":
		return Location{"Termez", "Asia/Samarkand"}
	case "TML":
		return Location{"Tamale", "Africa/Accra"}
	case "TMM":
		return Location{"", "Indian/Antananarivo"}
	case "TMP":
		return Location{"Tampere / Pirkkala", "Europe/Helsinki"}
	case "TMR":
		return Location{"Tamanrasset", "Africa/Algiers"}
	case "TMS":
		return Location{"Sao Tome", "Africa/Sao_Tome"}
	case "TMT":
		return Location{"Oriximina", "America/Santarem"}
	case "TMW":
		return Location{"Tamworth", "Australia/Sydney"}
	case "TMX":
		return Location{"Timimoun", "Africa/Algiers"}
	case "TNA":
		return Location{"Jinan", "Asia/Shanghai"}
	case "TNC":
		return Location{"Tin City", "America/Nome"}
	case "TNE":
		return Location{"", "Asia/Tokyo"}
	case "TNG":
		return Location{"Tangier", "Africa/Casablanca"}
	case "TNH":
		return Location{"Tonghua", "Asia/Shanghai"}
	case "TNJ":
		return Location{"Tanjung Pinang-Bintan Island", "Asia/Jakarta"}
	case "TNN":
		return Location{"Tainan City", "Asia/Taipei"}
	case "TNO":
		return Location{"Santa Cruz", "America/Costa_Rica"}
	case "TNR":
		return Location{"Antananarivo", "Indian/Antananarivo"}
	case "TOB":
		return Location{"Tobruk", "Africa/Tripoli"}
	case "TOE":
		return Location{"Tozeur", "Africa/Tunis"}
	case "TOF":
		return Location{"Tomsk", "Asia/Tomsk"}
	case "TOG":
		return Location{"Togiak Village", "America/Anchorage"}
	case "TOH":
		return Location{"Loh/Linua", "Pacific/Efate"}
	case "TOL":
		return Location{"Toledo", "America/New_York"}
	case "TOS":
		return Location{"Tromso", "Europe/Oslo"}
	case "TOY":
		return Location{"Toyama", "Asia/Tokyo"}
	case "TPA":
		return Location{"Tampa", "America/New_York"}
	case "TPE":
		return Location{"Taipei", "Asia/Taipei"}
	case "TPP":
		return Location{"Tarapoto", "America/Lima"}
	case "TPQ":
		return Location{"Tepic", "America/Mazatlan"}
	case "TPS":
		return Location{"Trapani", "Europe/Rome"}
	case "TRC":
		return Location{"Torreon", "America/Monterrey"}
	case "TRD":
		return Location{"Trondheim", "Europe/Oslo"}
	case "TRE":
		return Location{"Balemartine", "Europe/London"}
	case "TRF":
		return Location{"Torp", "Europe/Oslo"}
	case "TRG":
		return Location{"Tauranga", "Pacific/Auckland"}
	case "TRI":
		return Location{"Bristol/Johnson/Kingsport", "America/New_York"}
	case "TRK":
		return Location{"Tarakan Island", "Asia/Makassar"}
	case "TRN":
		return Location{"Torino", "Europe/Rome"}
	case "TRR":
		return Location{"Trincomalee", "Asia/Colombo"}
	case "TRS":
		return Location{"Trieste", "Europe/Rome"}
	case "TRU":
		return Location{"Trujillo", "America/Lima"}
	case "TRV":
		return Location{"Trivandrum", "Asia/Kolkata"}
	case "TRW":
		return Location{"Tarawa", "Pacific/Tarawa"}
	case "TRZ":
		return Location{"Tiruchirappally", "Asia/Kolkata"}
	case "TSA":
		return Location{"Taipei City", "Asia/Taipei"}
	case "TSF":
		return Location{"Treviso", "Europe/Rome"}
	case "TSJ":
		return Location{"Tsushima", "Asia/Tokyo"}
	case "TSM":
		return Location{"Taos", "America/Denver"}
	case "TSN":
		return Location{"Tianjin", "Asia/Shanghai"}
	case "TSR":
		return Location{"Timisoara", "Europe/Bucharest"}
	case "TST":
		return Location{"", "Asia/Bangkok"}
	case "TSV":
		return Location{"Townsville", "Australia/Brisbane"}
	case "TTA":
		return Location{"Tan Tan", "Africa/Casablanca"}
	case "TTE":
		return Location{"Sango-Ternate Island", "Asia/Jayapura"}
	case "TTJ":
		return Location{"Tottori", "Asia/Tokyo"}
	case "TTN":
		return Location{"Trenton", "America/New_York"}
	case "TTQ":
		return Location{"Roxana", "America/Costa_Rica"}
	case "TTT":
		return Location{"Taitung City", "Asia/Taipei"}
	case "TTU":
		return Location{"", "Africa/Casablanca"}
	case "TUB":
		return Location{"", "Pacific/Tahiti"}
	case "TUC":
		return Location{"San Miguel de Tucuman", "America/Argentina/Tucuman"}
	case "TUF":
		return Location{"Tours/Val de Loire (Loire Valley)", "Europe/Paris"}
	case "TUG":
		return Location{"Tuguegarao City", "Asia/Manila"}
	case "TUI":
		return Location{"", "Asia/Riyadh"}
	case "TUK":
		return Location{"Turbat", "Asia/Karachi"}
	case "TUL":
		return Location{"Tulsa", "America/Chicago"}
	case "TUN":
		return Location{"Tunis", "Africa/Tunis"}
	case "TUO":
		return Location{"Taupo", "Pacific/Auckland"}
	case "TUP":
		return Location{"Tupelo", "America/Chicago"}
	case "TUR":
		return Location{"Tucurui", "America/Belem"}
	case "TUS":
		return Location{"Tucson", "America/Phoenix"}
	case "TUU":
		return Location{"", "Asia/Riyadh"}
	case "TVC":
		return Location{"Traverse City", "America/Detroit"}
	case "TVF":
		return Location{"Thief River Falls", "America/Chicago"}
	case "TVU":
		return Location{"Matei", "Pacific/Fiji"}
	case "TVY":
		return Location{"Dawei", "Asia/Yangon"}
	case "TWF":
		return Location{"Twin Falls", "America/Boise"}
	case "TWU":
		return Location{"Tawau", "Asia/Kuching"}
	case "TXF":
		return Location{"Teixeira De Freitas", "America/Bahia"}
	case "TXK":
		return Location{"Texarkana", "America/Chicago"}
	case "TXN":
		return Location{"Huangshan", "Asia/Shanghai"}
	case "TYD":
		return Location{"Tynda", "Asia/Yakutsk"}
	case "TYF":
		return Location{"", "Europe/Stockholm"}
	case "TYL":
		return Location{"", "America/Lima"}
	case "TYN":
		return Location{"Taiyuan", "Asia/Shanghai"}
	case "TYR":
		return Location{"Tyler", "America/Chicago"}
	case "TYS":
		return Location{"Knoxville", "America/New_York"}
	case "TZL":
		return Location{"Tuzla", "Europe/Sarajevo"}
	case "TZX":
		return Location{"Trabzon", "Europe/Istanbul"}
	case "UAH":
		return Location{"Ua Huka", "Pacific/Marquesas"}
	case "UAK":
		return Location{"Narsarsuaq", "America/Nuuk"}
	case "UAP":
		return Location{"Ua Pou", "Pacific/Marquesas"}
	case "UAQ":
		return Location{"San Juan", "America/Argentina/San_Juan"}
	case "UAS":
		return Location{"Samburu South", "Africa/Nairobi"}
	case "UBA":
		return Location{"Uberaba", "America/Sao_Paulo"}
	case "UBB":
		return Location{"Mabuiag Island", "Australia/Brisbane"}
	case "UBJ":
		return Location{"Ube", "Asia/Tokyo"}
	case "UBN":
		return Location{"Ulaanbaatar", "Asia/Ulaanbaatar"}
	case "UBP":
		return Location{"Ubon Ratchathani", "Asia/Bangkok"}
	case "UCT":
		return Location{"Ukhta", "Europe/Moscow"}
	case "UDI":
		return Location{"Uberlandia", "America/Sao_Paulo"}
	case "UDR":
		return Location{"Udaipur", "Asia/Kolkata"}
	case "UEL":
		return Location{"Quelimane", "Africa/Maputo"}
	case "UEO":
		return Location{"", "Asia/Tokyo"}
	case "UET":
		return Location{"Quetta", "Asia/Karachi"}
	case "UFA":
		return Location{"Ufa", "Asia/Yekaterinburg"}
	case "UGC":
		return Location{"Urgench", "Asia/Samarkand"}
	case "UIB":
		return Location{"Quibdo", "America/Bogota"}
	case "UIH":
		return Location{"Quy Nohn", "Asia/Ho_Chi_Minh"}
	case "UII":
		return Location{"Utila Island", "America/Tegucigalpa"}
	case "UIN":
		return Location{"Quincy", "America/Chicago"}
	case "UIO":
		return Location{"Quito", "America/Guayaquil"}
	case "UKA":
		return Location{"Ukunda", "Africa/Nairobi"}
	case "UKB":
		return Location{"Kobe", "Asia/Tokyo"}
	case "UKG":
		return Location{"Ust-Kuyga", "Asia/Vladivostok"}
	case "UKK":
		return Location{"Ust Kamenogorsk", "Asia/Almaty"}
	case "UKX":
		return Location{"Ust-Kut", "Asia/Irkutsk"}
	case "ULB":
		return Location{"Ambryn Island", "Pacific/Efate"}
	case "ULG":
		return Location{"", "Asia/Hovd"}
	case "ULH":
		return Location{"Al-'Ula", "Asia/Riyadh"}
	case "ULK":
		return Location{"Lensk", "Asia/Yakutsk"}
	case "ULO":
		return Location{"", "Asia/Hovd"}
	case "ULP":
		return Location{"", "Australia/Brisbane"}
	case "ULV":
		return Location{"Ulyanovsk", "Europe/Ulyanovsk"}
	case "UME":
		return Location{"Umea", "Europe/Stockholm"}
	case "UMU":
		return Location{"Umuarama", "America/Sao_Paulo"}
	case "UNA":
		return Location{"Una", "America/Bahia"}
	case "UNG":
		return Location{"Kiunga", "Pacific/Port_Moresby"}
	case "UNK":
		return Location{"Unalakleet", "America/Anchorage"}
	case "UNN":
		return Location{"", "Asia/Bangkok"}
	case "UPG":
		return Location{"Ujung Pandang-Celebes Island", "Asia/Makassar"}
	case "UPN":
		return Location{"", "America/Mexico_City"}
	case "URA":
		return Location{"Uralsk", "Asia/Oral"}
	case "URC":
		return Location{"Urumqi", "Asia/Shanghai"}
	case "URE":
		return Location{"Kuressaare", "Europe/Tallinn"}
	case "URG":
		return Location{"Uruguaiana", "America/Sao_Paulo"}
	case "URJ":
		return Location{"Uray", "Asia/Yekaterinburg"}
	case "URT":
		return Location{"Surat Thani", "Asia/Bangkok"}
	case "URY":
		return Location{"", "Asia/Riyadh"}
	case "USH":
		return Location{"Ushuahia", "America/Argentina/Ushuaia"}
	case "USK":
		return Location{"Usinsk", "Europe/Moscow"}
	case "USM":
		return Location{"Na Thon (Ko Samui Island)", "Asia/Bangkok"}
	case "USN":
		return Location{"Ulsan", "Asia/Seoul"}
	case "USR":
		return Location{"Ust-Nera", "Asia/Ust-Nera"}
	case "USU":
		return Location{"Coron", "Asia/Manila"}
	case "UTH":
		return Location{"Udon Thani", "Asia/Bangkok"}
	case "UTN":
		return Location{"Upington", "Africa/Johannesburg"}
	case "UTP":
		return Location{"Rayong", "Asia/Bangkok"}
	case "UTT":
		return Location{"Mthatha", "Africa/Johannesburg"}
	case "UUA":
		return Location{"Bugulma", "Europe/Moscow"}
	case "UUD":
		return Location{"Ulan Ude", "Asia/Irkutsk"}
	case "UUS":
		return Location{"Yuzhno-Sakhalinsk", "Asia/Sakhalin"}
	case "UVE":
		return Location{"Ouvea", "Pacific/Noumea"}
	case "UVF":
		return Location{"Vieux Fort", "America/St_Lucia"}
	case "UYN":
		return Location{"Yulin", "Asia/Shanghai"}
	case "UYU":
		return Location{"Quijarro", "America/La_Paz"}
	case "VAA":
		return Location{"Vaasa", "Europe/Helsinki"}
	case "VAG":
		return Location{"Varginha", "America/Sao_Paulo"}
	case "VAI":
		return Location{"", "Pacific/Port_Moresby"}
	case "VAK":
		return Location{"Chevak", "America/Nome"}
	case "VAM":
		return Location{"Maamigili", "Indian/Maldives"}
	case "VAN":
		return Location{"Van", "Europe/Istanbul"}
	case "VAO":
		return Location{"Suavanao", "Pacific/Guadalcanal"}
	case "VAR":
		return Location{"Varna", "Europe/Sofia"}
	case "VAS":
		return Location{"Sivas", "Europe/Istanbul"}
	case "VAV":
		return Location{"Vava'u Island", "Pacific/Tongatapu"}
	case "VAW":
		return Location{"Vardo", "Europe/Oslo"}
	case "VBA":
		return Location{"Aeng", "Asia/Yangon"}
	case "VBP":
		return Location{"Bokpyinn", "Asia/Yangon"}
	case "VBV":
		return Location{"Vanua Balavu", "Pacific/Fiji"}
	case "VBY":
		return Location{"Visby", "Europe/Stockholm"}
	case "VCA":
		return Location{"Can Tho", "Asia/Ho_Chi_Minh"}
	case "VCE":
		return Location{"Venezia", "Europe/Rome"}
	case "VCL":
		return Location{"Dung Quat Bay", "Asia/Ho_Chi_Minh"}
	case "VCP":
		return Location{"Campinas", "America/Sao_Paulo"}
	case "VCS":
		return Location{"Con Ong", "Asia/Ho_Chi_Minh"}
	case "VCT":
		return Location{"Victoria", "America/Chicago"}
	case "VDC":
		return Location{"Vitoria Da Conquista", "America/Bahia"}
	case "VDE":
		return Location{"El Hierro Island", "Atlantic/Canary"}
	case "VDH":
		return Location{"Dong Hoi", "Asia/Bangkok"}
	case "VDM":
		return Location{"Viedma / Carmen de Patagones", "America/Argentina/Salta"}
	case "VDS":
		return Location{"Vadso", "Europe/Oslo"}
	case "VDY":
		return Location{"", "Asia/Kolkata"}
	case "VDZ":
		return Location{"Valdez", "America/Anchorage"}
	case "VEE":
		return Location{"Venetie", "America/Anchorage"}
	case "VEL":
		return Location{"Vernal", "America/Denver"}
	case "VER":
		return Location{"Veracruz", "America/Mexico_City"}
	case "VFA":
		return Location{"Victoria Falls", "Africa/Harare"}
	case "VGA":
		return Location{"", "Asia/Kolkata"}
	case "VGO":
		return Location{"Vigo", "Europe/Madrid"}
	case "VGZ":
		return Location{"Villa Garzon", "America/Bogota"}
	case "VHC":
		return Location{"Saurimo", "Africa/Luanda"}
	case "VHM":
		return Location{"", "Europe/Stockholm"}
	case "VHZ":
		return Location{"Vahitahi", "Pacific/Tahiti"}
	case "VIE":
		return Location{"Vienna", "Europe/Vienna"}
	case "VIG":
		return Location{"El Vigia", "America/Caracas"}
	case "VII":
		return Location{"Vinh", "Asia/Bangkok"}
	case "VIJ":
		return Location{"Spanish Town", "America/Tortola"}
	case "VIL":
		return Location{"Dakhla", "Africa/El_Aaiun"}
	case "VIT":
		return Location{"Alava", "Europe/Madrid"}
	case "VIX":
		return Location{"Vitoria", "America/Sao_Paulo"}
	case "VJB":
		return Location{"Xai-Xai", "Africa/Maputo"}
	case "VKG":
		return Location{"Rach Gia", "Asia/Ho_Chi_Minh"}
	case "VKO":
		return Location{"Moscow", "Europe/Moscow"}
	case "VKT":
		return Location{"Vorkuta", "Europe/Moscow"}
	case "VLC":
		return Location{"Valencia", "Europe/Madrid"}
	case "VLD":
		return Location{"Valdosta", "America/New_York"}
	case "VLI":
		return Location{"Port Vila", "Pacific/Efate"}
	case "VLL":
		return Location{"Valladolid", "Europe/Madrid"}
	case "VLN":
		return Location{"Valencia", "America/Caracas"}
	case "VLS":
		return Location{"Valesdir", "Pacific/Efate"}
	case "VLV":
		return Location{"Valera", "America/Caracas"}
	case "VNO":
		return Location{"Vilnius", "Europe/Vilnius"}
	case "VNS":
		return Location{"Varanasi", "Asia/Kolkata"}
	case "VNX":
		return Location{"Vilanculo", "Africa/Maputo"}
	case "VNY":
		return Location{"Van Nuys", "America/Los_Angeles"}
	case "VOG":
		return Location{"Volgograd", "Europe/Volgograd"}
	case "VOL":
		return Location{"Nea Anchialos", "Europe/Athens"}
	case "VPE":
		return Location{"Ngiva", "Africa/Luanda"}
	case "VPS":
		return Location{"Valparaiso", "America/Chicago"}
	case "VPY":
		return Location{"Chimoio", "Africa/Maputo"}
	case "VQS":
		return Location{"Vieques Island", "America/Puerto_Rico"}
	case "VRA":
		return Location{"Varadero", "America/Havana"}
	case "VRB":
		return Location{"Vero Beach", "America/New_York"}
	case "VRC":
		return Location{"Virac", "Asia/Manila"}
	case "VRN":
		return Location{"Verona", "Europe/Rome"}
	case "VSA":
		return Location{"Villahermosa", "America/Mexico_City"}
	case "VST":
		return Location{"Stockholm / Vasteras", "Europe/Stockholm"}
	case "VTE":
		return Location{"Vientiane", "Asia/Vientiane"}
	case "VTZ":
		return Location{"Visakhapatnam", "Asia/Kolkata"}
	case "VUP":
		return Location{"Valledupar", "America/Bogota"}
	case "VUS":
		return Location{"Velikiy Ustyug", "Europe/Moscow"}
	case "VVC":
		return Location{"Villavicencio", "America/Bogota"}
	case "VVI":
		return Location{"Santa Cruz", "America/La_Paz"}
	case "VVO":
		return Location{"Vladivostok", "Asia/Vladivostok"}
	case "VVZ":
		return Location{"Illizi", "Africa/Algiers"}
	case "VXC":
		return Location{"Lichinga", "Africa/Maputo"}
	case "VXE":
		return Location{"Sao Pedro", "Atlantic/Cape_Verde"}
	case "VXO":
		return Location{"Vaxjo", "Europe/Stockholm"}
	case "WAA":
		return Location{"Wales", "America/Nome"}
	case "WAE":
		return Location{"", "Asia/Riyadh"}
	case "WAG":
		return Location{"Wanganui", "Pacific/Auckland"}
	case "WAW":
		return Location{"Warsaw", "Europe/Warsaw"}
	case "WBM":
		return Location{"", "Pacific/Port_Moresby"}
	case "WBQ":
		return Location{"Beaver", "America/Anchorage"}
	case "WDH":
		return Location{"Windhoek", "Africa/Windhoek"}
	case "WEF":
		return Location{"Weifang", "Asia/Shanghai"}
	case "WEH":
		return Location{"Weihai", "Asia/Shanghai"}
	case "WEI":
		return Location{"Weipa", "Australia/Brisbane"}
	case "WGA":
		return Location{"Wagga Wagga", "Australia/Sydney"}
	case "WGE":
		return Location{"", "Australia/Sydney"}
	case "WGP":
		return Location{"Waingapu-Sumba Island", "Asia/Makassar"}
	case "WHK":
		return Location{"", "Pacific/Auckland"}
	case "WIC":
		return Location{"Wick", "Europe/London"}
	case "WIL":
		return Location{"Nairobi", "Africa/Nairobi"}
	case "WIN":
		return Location{"", "Australia/Brisbane"}
	case "WJU":
		return Location{"Wonju", "Asia/Seoul"}
	case "WKA":
		return Location{"", "Pacific/Auckland"}
	case "WKJ":
		return Location{"Wakkanai", "Asia/Tokyo"}
	case "WLE":
		return Location{"", "Australia/Brisbane"}
	case "WLG":
		return Location{"Wellington", "Pacific/Auckland"}
	case "WLH":
		return Location{"Walaha", "Pacific/Efate"}
	case "WLK":
		return Location{"Selawik", "America/Anchorage"}
	case "WLS":
		return Location{"Wallis Island", "Pacific/Wallis"}
	case "WMI":
		return Location{"Warsaw", "Europe/Warsaw"}
	case "WMN":
		return Location{"", "Indian/Antananarivo"}
	case "WMO":
		return Location{"White Mountain", "America/Nome"}
	case "WMX":
		return Location{"Wamena-Papua Island", "Asia/Jayapura"}
	case "WNA":
		return Location{"Napakiak", "America/Anchorage"}
	case "WNP":
		return Location{"Naga", "Asia/Manila"}
	case "WNR":
		return Location{"", "Australia/Brisbane"}
	case "WNZ":
		return Location{"Wenzhou", "Asia/Shanghai"}
	case "WOL":
		return Location{"", "Australia/Sydney"}
	case "WPR":
		return Location{"Porvenir", "America/Punta_Arenas"}
	case "WPU":
		return Location{"Puerto Williams", "America/Argentina/Ushuaia"}
	case "WRE":
		return Location{"", "Pacific/Auckland"}
	case "WRG":
		return Location{"Wrangell", "America/Sitka"}
	case "WRO":
		return Location{"Wroclaw", "Europe/Warsaw"}
	case "WRZ":
		return Location{"Weerawila", "Asia/Colombo"}
	case "WSN":
		return Location{"South Naknek", "America/Anchorage"}
	case "WSZ":
		return Location{"", "Pacific/Auckland"}
	case "WTB":
		return Location{"Wellcamp", "Australia/Brisbane"}
	case "WTK":
		return Location{"Noatak", "America/Nome"}
	case "WUH":
		return Location{"Wuhan", "Asia/Shanghai"}
	case "WUN":
		return Location{"", "Australia/Perth"}
	case "WUS":
		return Location{"Wuyishan", "Asia/Shanghai"}
	case "WUX":
		return Location{"Wuxi", "Asia/Shanghai"}
	case "WUZ":
		return Location{"Wuzhou", "Asia/Shanghai"}
	case "WVB":
		return Location{"Walvis Bay", "Africa/Windhoek"}
	case "WWK":
		return Location{"Wewak", "Pacific/Port_Moresby"}
	case "WWT":
		return Location{"Newtok", "America/Nome"}
	case "WXN":
		return Location{"Wanxian", "Asia/Shanghai"}
	case "WYA":
		return Location{"Whyalla", "Australia/Adelaide"}
	case "WYS":
		return Location{"West Yellowstone", "America/Denver"}
	case "XAP":
		return Location{"Chapeco", "America/Sao_Paulo"}
	case "XBE":
		return Location{"Bearskin Lake", "America/Rainy_River"}
	case "XCH":
		return Location{"Christmas Island", "Indian/Christmas"}
	case "XCR":
		return Location{"Chalons/Vatry", "Europe/Paris"}
	case "XFN":
		return Location{"Xiangfan", "Asia/Shanghai"}
	case "XGR":
		return Location{"Kangiqsualujjuaq", "America/Toronto"}
	case "XIC":
		return Location{"Xichang", "Asia/Shanghai"}
	case "XIL":
		return Location{"Xilinhot", "Asia/Shanghai"}
	case "XIY":
		return Location{"Xianyang", "Asia/Shanghai"}
	case "XKH":
		return Location{"Xieng Khouang", "Asia/Vientiane"}
	case "XLS":
		return Location{"Saint Louis", "Africa/Dakar"}
	case "XMH":
		return Location{"", "Pacific/Tahiti"}
	case "XMN":
		return Location{"Xiamen", "Asia/Shanghai"}
	case "XMY":
		return Location{"Yam Island", "Australia/Brisbane"}
	case "XNA":
		return Location{"Fayetteville/Springdale/", "America/Chicago"}
	case "XNN":
		return Location{"Xining", "Asia/Shanghai"}
	case "XPL":
		return Location{"Comayagua", "America/Tegucigalpa"}
	case "XQP":
		return Location{"Quepos", "America/Costa_Rica"}
	case "XRY":
		return Location{"Jerez de la Forntera", "Europe/Madrid"}
	case "XSC":
		return Location{"", "America/Grand_Turk"}
	case "XSP":
		return Location{"Seletar", "Asia/Kuala_Lumpur"}
	case "XTG":
		return Location{"", "Australia/Brisbane"}
	case "XUZ":
		return Location{"Xuzhou", "Asia/Shanghai"}
	case "YAA":
		return Location{"Anahim Lake", "America/Vancouver"}
	case "YAB":
		return Location{"", "America/Rankin_Inlet"}
	case "YAC":
		return Location{"Cat Lake", "America/Rainy_River"}
	case "YAG":
		return Location{"Fort Frances", "America/Rainy_River"}
	case "YAK":
		return Location{"Yakutat", "America/Yakutat"}
	case "YAM":
		return Location{"Sault Ste Marie", "America/Detroit"}
	case "YAP":
		return Location{"Yap Island", "Pacific/Chuuk"}
	case "YAT":
		return Location{"Attawapiskat", "America/Nipigon"}
	case "YAY":
		return Location{"St. Anthony", "America/St_Johns"}
	case "YAZ":
		return Location{"Tofino", "America/Vancouver"}
	case "YBB":
		return Location{"Kugaaruk", "America/Cambridge_Bay"}
	case "YBE":
		return Location{"Uranium City", "America/Regina"}
	case "YBG":
		return Location{"Bagotville", "America/Toronto"}
	case "YBK":
		return Location{"Baker Lake", "America/Rankin_Inlet"}
	case "YBL":
		return Location{"Campbell River", "America/Vancouver"}
	case "YBP":
		return Location{"Yibin", "Asia/Shanghai"}
	case "YBR":
		return Location{"Brandon", "America/Winnipeg"}
	case "YBX":
		return Location{"Lourdes-De-Blanc-Sablon", "America/Blanc-Sablon"}
	case "YCB":
		return Location{"Cambridge Bay", "America/Cambridge_Bay"}
	case "YCD":
		return Location{"Nanaimo", "America/Vancouver"}
	case "YCG":
		return Location{"Castlegar", "America/Vancouver"}
	case "YCK":
		return Location{"Colville Lake", "America/Inuvik"}
	case "YCO":
		return Location{"Kugluktuk", "America/Cambridge_Bay"}
	case "YCS":
		return Location{"Chesterfield Inlet", "America/Rankin_Inlet"}
	case "YCU":
		return Location{"Yuncheng", "Asia/Shanghai"}
	case "YCY":
		return Location{"Clyde River", "America/Iqaluit"}
	case "YDA":
		return Location{"Dawson City", "America/Dawson"}
	case "YDF":
		return Location{"Deer Lake", "America/St_Johns"}
	case "YDP":
		return Location{"Nain", "America/Goose_Bay"}
	case "YEG":
		return Location{"Edmonton", "America/Edmonton"}
	case "YEI":
		return Location{"Bursa", "Europe/Istanbul"}
	case "YEK":
		return Location{"Arviat", "America/Rankin_Inlet"}
	case "YEV":
		return Location{"Inuvik", "America/Inuvik"}
	case "YFA":
		return Location{"Fort Albany", "America/Nipigon"}
	case "YFB":
		return Location{"Iqaluit", "America/Iqaluit"}
	case "YFC":
		return Location{"Fredericton", "America/Moncton"}
	case "YFH":
		return Location{"Fort Hope", "America/Nipigon"}
	case "YFJ":
		return Location{"Wekweeti", "America/Yellowknife"}
	case "YFO":
		return Location{"Flin Flon", "America/Winnipeg"}
	case "YFS":
		return Location{"Fort Simpson", "America/Inuvik"}
	case "YFX":
		return Location{"St. Lewis", "America/St_Johns"}
	case "YGH":
		return Location{"Fort Good Hope", "America/Inuvik"}
	case "YGJ":
		return Location{"Yonago", "Asia/Tokyo"}
	case "YGK":
		return Location{"Kingston", "America/Toronto"}
	case "YGL":
		return Location{"La Grande Riviere", "America/Toronto"}
	case "YGP":
		return Location{"Gaspe", "America/Toronto"}
	case "YGR":
		return Location{"Iles-de-la-Madeleine", "America/Halifax"}
	case "YGT":
		return Location{"Igloolik", "America/Iqaluit"}
	case "YGW":
		return Location{"Kuujjuarapik", "America/Iqaluit"}
	case "YGX":
		return Location{"Gillam", "America/Winnipeg"}
	case "YGZ":
		return Location{"Grise Fiord", "America/Iqaluit"}
	case "YHA":
		return Location{"Port Hope Simpson", "America/St_Johns"}
	case "YHD":
		return Location{"Dryden", "America/Rainy_River"}
	case "YHI":
		return Location{"Ulukhaktok", "America/Yellowknife"}
	case "YHK":
		return Location{"Gjoa Haven", "America/Cambridge_Bay"}
	case "YHM":
		return Location{"Hamilton", "America/Toronto"}
	case "YHO":
		return Location{"Hopedale", "America/Goose_Bay"}
	case "YHP":
		return Location{"Poplar Hill", "America/Rainy_River"}
	case "YHR":
		return Location{"Chevery", "America/Blanc-Sablon"}
	case "YHU":
		return Location{"Montreal", "America/Toronto"}
	case "YHY":
		return Location{"Hay River", "America/Yellowknife"}
	case "YHZ":
		return Location{"Halifax", "America/Halifax"}
	case "YIA":
		return Location{"Yogyakarta", "Asia/Jakarta"}
	case "YIF":
		return Location{"St-Augustin", "America/Blanc-Sablon"}
	case "YIH":
		return Location{"Yichang", "Asia/Shanghai"}
	case "YIK":
		return Location{"Ivujivik", "America/Iqaluit"}
	case "YIN":
		return Location{"Yining", "Asia/Shanghai"}
	case "YIO":
		return Location{"Pond Inlet", "America/Iqaluit"}
	case "YIW":
		return Location{"Yiwu", "Asia/Shanghai"}
	case "YKA":
		return Location{"Kamloops", "America/Vancouver"}
	case "YKF":
		return Location{"Kitchener", "America/Toronto"}
	case "YKG":
		return Location{"Kangirsuk", "America/Toronto"}
	case "YKL":
		return Location{"Schefferville", "America/Toronto"}
	case "YKM":
		return Location{"Yakima", "America/Los_Angeles"}
	case "YKO":
		return Location{"Yuksekova", "Europe/Istanbul"}
	case "YKQ":
		return Location{"Waskaganish", "America/Toronto"}
	case "YKS":
		return Location{"Yakutsk", "Asia/Yakutsk"}
	case "YKU":
		return Location{"Chisasibi", "America/Toronto"}
	case "YLC":
		return Location{"Kimmirut", "America/Iqaluit"}
	case "YLE":
		return Location{"Whati", "America/Yellowknife"}
	case "YLH":
		return Location{"Lansdowne House", "America/Nipigon"}
	case "YLL":
		return Location{"Lloydminster", "America/Edmonton"}
	case "YLW":
		return Location{"Kelowna", "America/Vancouver"}
	case "YMH":
		return Location{"Mary's Harbour", "America/St_Johns"}
	case "YMM":
		return Location{"Fort McMurray", "America/Edmonton"}
	case "YMN":
		return Location{"Makkovik", "America/Goose_Bay"}
	case "YMO":
		return Location{"Moosonee", "America/Nipigon"}
	case "YMT":
		return Location{"Chibougamau", "America/Toronto"}
	case "YNA":
		return Location{"Natashquan", "America/Toronto"}
	case "YNB":
		return Location{"", "Asia/Riyadh"}
	case "YNC":
		return Location{"Wemindji", "America/Toronto"}
	case "YNJ":
		return Location{"Yanji", "Asia/Shanghai"}
	case "YNL":
		return Location{"Points North Landing", "America/Regina"}
	case "YNO":
		return Location{"North Spirit Lake", "America/Rainy_River"}
	case "YNS":
		return Location{"Nemiscau", "America/Toronto"}
	case "YNT":
		return Location{"Yantai", "Asia/Shanghai"}
	case "YNY":
		return Location{"Sokcho / Gangneung", "Asia/Seoul"}
	case "YNZ":
		return Location{"Yancheng", "Asia/Shanghai"}
	case "YOC":
		return Location{"Old Crow", "America/Dawson"}
	case "YOG":
		return Location{"Ogoki Post", "America/Nipigon"}
	case "YOJ":
		return Location{"High Level", "America/Edmonton"}
	case "YOL":
		return Location{"Yola", "Africa/Lagos"}
	case "YOW":
		return Location{"Ottawa", "America/Toronto"}
	case "YPA":
		return Location{"Prince Albert", "America/Regina"}
	case "YPH":
		return Location{"Inukjuak", "America/Toronto"}
	case "YPJ":
		return Location{"Aupaluk", "America/Toronto"}
	case "YPM":
		return Location{"Pikangikum", "America/Rainy_River"}
	case "YPO":
		return Location{"Peawanuck", "America/Nipigon"}
	case "YPR":
		return Location{"Prince Rupert", "America/Vancouver"}
	case "YPW":
		return Location{"Powell River", "America/Vancouver"}
	case "YPX":
		return Location{"Puvirnituq", "America/Toronto"}
	case "YPY":
		return Location{"Fort Chipewyan", "America/Edmonton"}
	case "YQB":
		return Location{"Quebec", "America/Toronto"}
	case "YQC":
		return Location{"Quaqtaq", "America/Iqaluit"}
	case "YQD":
		return Location{"The Pas", "America/Winnipeg"}
	case "YQG":
		return Location{"Windsor", "America/Toronto"}
	case "YQK":
		return Location{"Kenora", "America/Rainy_River"}
	case "YQL":
		return Location{"Lethbridge", "America/Edmonton"}
	case "YQM":
		return Location{"Moncton", "America/Moncton"}
	case "YQQ":
		return Location{"Comox", "America/Vancouver"}
	case "YQR":
		return Location{"Regina", "America/Regina"}
	case "YQT":
		return Location{"Thunder Bay", "America/Thunder_Bay"}
	case "YQU":
		return Location{"Grande Prairie", "America/Edmonton"}
	case "YQX":
		return Location{"Gander", "America/St_Johns"}
	case "YQY":
		return Location{"Sydney", "America/Glace_Bay"}
	case "YQZ":
		return Location{"Quesnel", "America/Vancouver"}
	case "YRA":
		return Location{"Gameti", "America/Yellowknife"}
	case "YRB":
		return Location{"Resolute Bay", "America/Resolute"}
	case "YRF":
		return Location{"Cartwright", "America/Goose_Bay"}
	case "YRG":
		return Location{"Rigolet", "America/Goose_Bay"}
	case "YRL":
		return Location{"Red Lake", "America/Rainy_River"}
	case "YRT":
		return Location{"Rankin Inlet", "America/Rankin_Inlet"}
	case "YSB":
		return Location{"Sudbury", "America/Toronto"}
	case "YSF":
		return Location{"Stony Rapids", "America/Regina"}
	case "YSG":
		return Location{"Lutselk'e", "America/Yellowknife"}
	case "YSJ":
		return Location{"Saint John", "America/Moncton"}
	case "YSK":
		return Location{"Sanikiluaq", "America/Iqaluit"}
	case "YSM":
		return Location{"Fort Smith", "America/Edmonton"}
	case "YTE":
		return Location{"Cape Dorset", "America/Iqaluit"}
	case "YTH":
		return Location{"Thompson", "America/Winnipeg"}
	case "YTQ":
		return Location{"Tasiujaq", "America/Toronto"}
	case "YTS":
		return Location{"Timmins", "America/Toronto"}
	case "YTZ":
		return Location{"Toronto", "America/Toronto"}
	case "YUD":
		return Location{"Umiujaq", "America/Iqaluit"}
	case "YUL":
		return Location{"Montreal", "America/Toronto"}
	case "YUM":
		return Location{"Yuma", "America/Phoenix"}
	case "YUS":
		return Location{"Yushu", "Asia/Shanghai"}
	case "YUT":
		return Location{"Repulse Bay", "America/Rankin_Inlet"}
	case "YUX":
		return Location{"Hall Beach", "America/Iqaluit"}
	case "YUY":
		return Location{"Rouyn-Noranda", "America/Toronto"}
	case "YVB":
		return Location{"Bonaventure", "America/Toronto"}
	case "YVC":
		return Location{"La Ronge", "America/Regina"}
	case "YVM":
		return Location{"Qikiqtarjuaq", "America/Pangnirtung"}
	case "YVO":
		return Location{"Val-d'Or", "America/Toronto"}
	case "YVP":
		return Location{"Kuujjuaq", "America/Toronto"}
	case "YVQ":
		return Location{"Norman Wells", "America/Inuvik"}
	case "YVR":
		return Location{"Vancouver", "America/Vancouver"}
	case "YVZ":
		return Location{"Deer Lake", "America/Rainy_River"}
	case "YWB":
		return Location{"Kangiqsujuaq", "America/Toronto"}
	case "YWG":
		return Location{"Winnipeg", "America/Winnipeg"}
	case "YWJ":
		return Location{"Deline", "America/Inuvik"}
	case "YWK":
		return Location{"Wabush", "America/Goose_Bay"}
	case "YWL":
		return Location{"Williams Lake", "America/Vancouver"}
	case "YWP":
		return Location{"Webequie", "America/Nipigon"}
	case "YXC":
		return Location{"Cranbrook", "America/Edmonton"}
	case "YXE":
		return Location{"Saskatoon", "America/Regina"}
	case "YXH":
		return Location{"Medicine Hat", "America/Edmonton"}
	case "YXJ":
		return Location{"Fort St.John", "America/Dawson_Creek"}
	case "YXL":
		return Location{"Sioux Lookout", "America/Rainy_River"}
	case "YXN":
		return Location{"Whale Cove", "America/Rankin_Inlet"}
	case "YXP":
		return Location{"Pangnirtung", "America/Pangnirtung"}
	case "YXS":
		return Location{"Prince George", "America/Vancouver"}
	case "YXT":
		return Location{"Terrace", "America/Vancouver"}
	case "YXU":
		return Location{"London", "America/Toronto"}
	case "YXX":
		return Location{"Abbotsford", "America/Los_Angeles"}
	case "YXY":
		return Location{"Whitehorse", "America/Whitehorse"}
	case "YYB":
		return Location{"North Bay", "America/Toronto"}
	case "YYC":
		return Location{"Calgary", "America/Edmonton"}
	case "YYD":
		return Location{"Smithers", "America/Vancouver"}
	case "YYE":
		return Location{"Fort Nelson", "America/Fort_Nelson"}
	case "YYF":
		return Location{"Penticton", "America/Vancouver"}
	case "YYG":
		return Location{"Charlottetown", "America/Halifax"}
	case "YYH":
		return Location{"Taloyoak", "America/Cambridge_Bay"}
	case "YYJ":
		return Location{"Victoria", "America/Vancouver"}
	case "YYQ":
		return Location{"Churchill", "America/Winnipeg"}
	case "YYR":
		return Location{"Goose Bay", "America/Goose_Bay"}
	case "YYT":
		return Location{"St. John's", "America/St_Johns"}
	case "YYY":
		return Location{"Mont-Joli", "America/Toronto"}
	case "YYZ":
		return Location{"Toronto", "America/Toronto"}
	case "YZF":
		return Location{"Yellowknife", "America/Yellowknife"}
	case "YZG":
		return Location{"Salluit", "America/Toronto"}
	case "YZP":
		return Location{"Sandspit", "America/Vancouver"}
	case "YZS":
		return Location{"Coral Harbour", "America/Atikokan"}
	case "YZT":
		return Location{"Port Hardy", "America/Vancouver"}
	case "YZV":
		return Location{"Sept-Iles", "America/Toronto"}
	case "YZY":
		return Location{"Mackenzie", "America/Vancouver"}
	case "YZZ":
		return Location{"Trail", "America/Vancouver"}
	case "ZAD":
		return Location{"Zadar", "Europe/Zagreb"}
	case "ZAG":
		return Location{"Zagreb", "Europe/Zagreb"}
	case "ZAH":
		return Location{"Zahedan", "Asia/Tehran"}
	case "ZAL":
		return Location{"Valdivia", "America/Santiago"}
	case "ZAM":
		return Location{"Zamboanga City", "Asia/Manila"}
	case "ZAT":
		return Location{"Zhaotong", "Asia/Shanghai"}
	case "ZAZ":
		return Location{"Zaragoza", "Europe/Madrid"}
	case "ZBF":
		return Location{"Bathurst", "America/Moncton"}
	case "ZBR":
		return Location{"Chabahar", "Asia/Tehran"}
	case "ZCL":
		return Location{"Zacatecas", "America/Mexico_City"}
	case "ZCO":
		return Location{"Temuco", "America/Santiago"}
	case "ZEL":
		return Location{"Bella Bella", "America/Vancouver"}
	case "ZEM":
		return Location{"Eastmain River", "America/Toronto"}
	case "ZER":
		return Location{"", "Asia/Kolkata"}
	case "ZFD":
		return Location{"Fond-Du-Lac", "America/Regina"}
	case "ZFN":
		return Location{"Tulita", "America/Inuvik"}
	case "ZGU":
		return Location{"Gaua Island", "Pacific/Efate"}
	case "ZHA":
		return Location{"Zhanjiang", "Asia/Shanghai"}
	case "ZHY":
		return Location{"Zhongwei", "Asia/Shanghai"}
	case "ZIA":
		return Location{"Trento", "Europe/Rome"}
	case "ZIG":
		return Location{"Ziguinchor", "Africa/Dakar"}
	case "ZIH":
		return Location{"Ixtapa", "America/Mexico_City"}
	case "ZIX":
		return Location{"Zhigansk", "Asia/Yakutsk"}
	case "ZKE":
		return Location{"Kashechewan", "America/Nipigon"}
	case "ZKP":
		return Location{"Kasompe", "Africa/Lusaka"}
	case "ZLO":
		return Location{"Manzanillo", "America/Mexico_City"}
	case "ZLT":
		return Location{"La Tabatiere", "America/Blanc-Sablon"}
	case "ZMT":
		return Location{"Masset", "America/Vancouver"}
	case "ZNE":
		return Location{"Newman", "Australia/Perth"}
	case "ZNZ":
		return Location{"Kiembi Samaki", "Africa/Dar_es_Salaam"}
	case "ZOS":
		return Location{"Osorno", "America/Santiago"}
	case "ZPB":
		return Location{"Sachigo Lake", "America/Rainy_River"}
	case "ZQN":
		return Location{"Queenstown", "Pacific/Auckland"}
	case "ZRH":
		return Location{"Zurich", "Europe/Zurich"}
	case "ZSA":
		return Location{"San Salvador", "America/Nassau"}
	case "ZSE":
		return Location{"St Pierre", "Indian/Reunion"}
	case "ZSJ":
		return Location{"Sandy Lake", "America/Rainy_River"}
	case "ZTA":
		return Location{"", "Pacific/Tahiti"}
	case "ZTB":
		return Location{"Tete-a-la-Baleine", "America/Blanc-Sablon"}
	case "ZTH":
		return Location{"Zakynthos Island", "Europe/Athens"}
	case "ZUH":
		return Location{"Zhuhai", "Asia/Shanghai"}
	case "ZUM":
		return Location{"Churchill Falls", "America/Goose_Bay"}
	case "ZVK":
		return Location{"", "Asia/Bangkok"}
	case "ZWL":
		return Location{"Wollaston Lake", "America/Regina"}
	case "ZYI":
		return Location{"Zunyi", "Asia/Shanghai"}
	case "ZYL":
		return Location{"Sylhet", "Asia/Dhaka"}
	}
	return Location{"Not supported IATA Code", "Not supported IATA Code"}
}
