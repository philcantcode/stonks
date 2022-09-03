<?php
$GLOBALS['root'] = $_SERVER['DOCUMENT_ROOT'];

// Remote server functions
$GLOBALS['api'] = "http://localhost:8008";
$GLOBALS['server'] = "http://server:8008";
$GLOBALS['web'] = "";

$GLOBALS['site-name'] = "stonks-web";
$GLOBALS['tab-name'] = "";
$GLOBALS['previous-dir'] = "";

function head($tab, $pdir) {
    $GLOBALS['tab-name'] = $tab;
    $GLOBALS['previous-dir'] = $pdir;

    include_once $GLOBALS['root'] . "/base/header.php";
}

function foot() {
    include_once $GLOBALS['root'] . "/base/footer.php";
}

function stockURL($id) {
    return $GLOBALS['web'] . "/tracking/stock.php?ticker=" . $id;
}

function searchByLabel($label, $array) {

    if (is_null($array)) {
        return array("Values" => array("Unknown " . $label));
    }

    foreach ($array as $row) {
        if (isset($row['Label']) && $row['Label'] == $label) {
            return $row;
        }
    }

    return array("Values" => array("Unknown " . $label));
}

// Loads a JSON file from the server root
function loadJSON($path, $isArray, $domain = "-1") {
    if ($domain == "-1")
    {
        $domain = $GLOBALS['server'];
    }

    $json = @file_get_contents($domain . $path);

    if ($json === false)
    {
        echo "TODO://REDIRECT TO 404";
        die();
    }

    return json_decode($json, $isArray);
}

function dataType($datType) {
    switch ($datType) 
    {
        case 0:
            return "EMPTY";
            break;
        case 1:
            return "IP";
            break;
        case 2:
            return "IP_RANGE";
            break;
        case 3:
            return "MAC";
            break;
        case 4:
            return "INTEGER";
            break;
        case 5:
            return "DECIMAL";
            break;
        case 6:
            return "BOOL";
            break;
        case 7:
            return "STRING";
            break;
        case 8:
            return "CIDR";
            break;
        case 9:
            return "IP6";
            break;
        case 10:
            return "MAC6";  
            break;
        case 11:
            return "CPE";  
            break; 
        case 12:
            return "CCI";  
            break;  
        case 13:
            return "CCBI";  
            break;
        case 14:
            return "IP_RANGE_LOW";  
            break; 
        case 15:
            return "IP_RANGE_HIGH";  
            break;     
        }

    return null;
}

?>