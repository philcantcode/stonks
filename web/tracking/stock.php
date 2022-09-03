<?php 
head(tab: "Stock", pdir: "Home");

$ticker = $_GET["ticker"];

$trackedStock = loadJSON("/get/tracking/ticker/" . $ticker, true);

$yahooStockInfo = loadJSON("/get/yahoo/basic/ticker-lookup/" . $ticker, true);

?>

<style>
    #tradingview_92d57 {
        height: 500px;
    }
</style>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 id="ticker-title" class="card-title">
                        <?php echo $ticker; ?>
                    </h5>
                    <form class="row g-2">
                        <div class="col-md-6">
                            <div class="input-group">
                                <span class="input-group-text">Company</span>
                                <input id="company-name" type="text" class="form-control details-auto-update" value="<?php echo $trackedStock["Company"]; ?>">
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="input-group">
                                <span class="input-group-text">Industry</span>
                                <input id="industry" type="text" class="form-control details-auto-update" value="<?php echo $trackedStock["Industry"]; ?>">
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="input-group">
                                <span class="input-group-text">Owned Stocks #</span>
                                <input id="owned-stock" type="text" class="form-control details-auto-update" value="<?php echo $trackedStock["OwnedStocks"]; ?>">
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="input-group">
                                <span class="input-group-text">Custom Group</span>
                                <input id="group" type="text" class="form-control details-auto-update" value="<?php echo $trackedStock["Group"]; ?>">
                            </div>
                        </div>
                        <textarea class="form-control details-auto-update" id="notes" rows="5" placeholder="Notes"><?php echo $trackedStock["Notes"]; ?></textarea>
                        <input type='button' class='btn btn-success float-right'value='Update'></input>
                    </form>

                    <h5 id="ticker-title" class="card-title">Technical Analysis Links</h5>
                    <a target="_blank" href="https://www.tradingview.com/symbols/<?php echo $ticker; ?>/financials-overview/">
                        <input type='button' class='btn btn-success float-right'value='Financials'></input>
                    </a>

                    <a target="_blank" href="https://www.sec.gov/edgar/search/">
                        <input type='button' class='btn btn-success float-right'value='SEC Filings'></input>
                    </a>

                    <a target="_blank" href="https://www.optionistics.com/quotes/option-prices/<?php echo $ticker; ?>">
                        <input type='button' class='btn btn-success float-right'value='Options'></input>
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Stock Chart</h5>
                    <div class="col-12" style="height:510px;">
                        <!-- TradingView Widget BEGIN -->
                        <div class="tradingview-widget-container">
                            <div id="tradingview_92d57"></div>
                            <div class="tradingview-widget-copyright"><a
                                    href="https://www.tradingview.com/symbols/NASDAQ-AAPL/" rel="noopener"
                                    target="_blank"><span class="blue-text">AAPL Chart</span></a> by TradingView</div>
                            <script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
                            <script type="text/javascript">
                                new TradingView.widget(
                                    {
                                        "autosize": true,
                                        "height": 500,
                                        "symbol": "<?php echo $ticker; ?>",
                                        "interval": "D",
                                        "timezone": "Etc/UTC",
                                        "theme": "light",
                                        "style": "1",
                                        "locale": "en",
                                        "toolbar_bg": "#f1f3f6",
                                        "enable_publishing": false,
                                        "allow_symbol_change": true,
                                        "container_id": "tradingview_92d57"
                                    }
                                );
                            </script>
                        </div>
                        <!-- TradingView Widget END -->
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Yahoo Finance Info</h5>

                    <table class="table datatable">
                        <thead>
                            <tr>
                                <th scope="col" data-sortable=""><a href="#" class="dataTable-sorter">Key</a></th>
                                <th scope="col" data-sortable=""><a href="#" class="dataTable-sorter">Value</a></th>
                            </tr>
                        </thead>
                        <tbody>
                            <?php

                            foreach($yahooStockInfo["data"] as $k => $v)
                            {
                                echo "
                                    <tr>
                                        <td>" . $k . "</td>
                                        <td>" . $v . "</td>
                                    </tr>
                                ";
                            }
                            ?>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Insider Trades</h5>

                    <div id="insider-trades-frame" class="row" style="height:600px;">
                        <!-- <iframe src="http://openinsider.com/search?q=<?php echo $ticker; ?> #results">

                        </iframe> -->
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Google Trends</h5>

                    <script type="text/javascript"
                        src="https://ssl.gstatic.com/trends_nrtr/3045_RC01/embed_loader.js"></script>
                    <script
                        type="text/javascript"> trends.embed.renderExploreWidget("TIMESERIES", { "comparisonItem": [{ "keyword": "<?php echo $ticker; ?> stock", "geo": "", "time": "2004-01-01 2022-08-29" }, { "keyword": "<?php echo $ticker; ?>", "geo": "", "time": "2004-01-01 2022-08-29" }], "category": 0, "property": "" }, { "exploreQuery": "date=all&q=<?php echo $ticker; ?>%20stock,<?php echo $ticker; ?>", "guestPath": "https://trends.google.com:443/trends/embed/" }); </script>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Options</h5>

                    <div id="insider-trades-frame" class="row" style="height:600px;">
                        <iframe src="https://www.optionistics.com/quotes/option-prices/<?php echo $ticker; ?> #mainbody">

                        </iframe>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<?php foot(); ?>

<script type="text/javascript">
$(window).on("load", function() 
{
    $titleEditID = -1;

    $(".details-auto-update").on('input', function()
    {
        if ($titleEditID != -1)
        {
            clearTimeout($titleEditID)
            console.log("cleared");
        }

        $titleEditID = setTimeout(function ()
        {
            $.ajax(
            {
                url: "<?php echo $GLOBALS['api']; ?>" + "/put/update/tracking",
                type: "POST",
                data: {
                    "ticker": "<?php echo $ticker; ?>",
                    "company": $("#company-name").val(),
                    "industry": $("#industry").val(),
                    "ownedStock": $("#owned-stock").val(),
                    "group": $("#group").val(),
                    "notes": $("#notes").val(),
                },
                success: function (response) 
                { 
                    console.log("Updated tracking");
                    $("#ticker-title").css('color', 'green');

                    setTimeout(function() 
                    {
                        $("#ticker-title").css('color', '#012970');
                    }, 2000);
                }
            });

            $titleEditID = -1;
        }, 1000); 
    });

    $.ajax(
    {
        url: "<?php echo $GLOBALS['api']; ?>" + "/get/open-insider/part",
        type: "POST",
        data: {
            "ticker": "<?php echo $ticker; ?>",
            "target": "table",
        },
        success: function (response) 
        { 
            console.log(response);
            $("#insider-trades-frame").html(response);
        }
    });
}); 
</script>