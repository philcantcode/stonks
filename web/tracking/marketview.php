<?php 
  head(tab: 'Market View', pdir: '.');

  $tracked = loadJSON('/get/tracking/all', true);
?>

<style>
    #tradingview_92d57 {
        height: 500px;
    }
</style>

<div class='row'>
    <div class='col-12'>
        <div class='row'>
            <h3 class='card-title'>Owned Stocks</h3>

            <?php
            foreach ($tracked as $stock)
            {
                if ($stock["OwnedStocks"] != "" && $stock["OwnedStocks"] > 0)
                {
                    echo "
                    <div class='col-3'>
                        <div class='card info-card sales-card'>
                            <div class='card-body'>
                                <h5 class='card-title'>" . $stock["Ticker"] . " <span>| " . $stock["OwnedStocks"] . "</span></h5>

                                <div class='d-flex align-items-center'>
                                    <div class='card-icon rounded-circle d-flex align-items-center justify-content-center'>
                                        <i class='bi bi-percent'></i>
                                    </div>
                                    <div class='ps-3'>
                                        <h6>
            
                                        </h6>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    ";
                }
            }
        ?>
        </div>
    </div>
</div>

<iframe src="https://www.youtube.com/embed/?listType=user_uploads&list=Fxempirenetwork" width="480" height="400"></iframe>


<div class='row'>
    <div class='col-12'>
        <div class='row'>
            <h3 class='card-title'>Industry ETFs</h3>

            <div class='col-12'>
                <div class='card info-card sales-card'>
                    <div class='card-body'>
                        <h5 class='card-title'>S&P 500<span></span></h5>

                        <div class="col-12" style="height:510px;">
                        <!-- TradingView Widget BEGIN -->
                        <div class="tradingview-widget-container">
                            <div id="tradingview_92d57"></div>
                            <div class="tradingview-widget-copyright"><a
                                    href="https://www.tradingview.com/symbols/SP:SPX/" rel="noopener"
                                    target="_blank">
                            <script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
                            <script type="text/javascript">
                                new TradingView.widget(
                                    {
                                        "autosize": true,
                                        "height": 500,
                                        "symbol": "SP:SPX",
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
</div>



<script type="text/javascript">
  $(window).on("load", function()
  {
    
  });
</script>

<?php foot(); ?>