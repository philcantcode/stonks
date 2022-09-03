<?php 
  head(tab: "Tracking", pdir: ".");

  $tracked = loadJSON("/get/tracking/all", true);
?>

<div class="col-12">
    <div class="row">
        <div class="col-12">
            <div class="card info-card sales-card">
                <div class="card-body">
                    <h5 class="card-title">Tracked Stock Tickers</h5>

                    <table class="table datatable" style="table-layout: fixed; word-wrap: break-word;">
                        <thead>
                            <tr>
                                <th scope="col" data-sortable=""><a href="#" class="dataTable-sorter">Ticker</a></th>
                                <th scope="col" data-sortable=""><a href="#" class="dataTable-sorter">Owned Stocks</a>
                                </th>
                                <th scope="col" data-sortable=""><a href="#" class="dataTable-sorter">Group</a></th>
                            </tr>
                        </thead>
                        <tbody>
                            <?php
                            foreach($tracked as $stock)
                            {
                                echo "
                                    <tr>
                                    <td><a href='" . stockURL($stock['Ticker']) . "'>" . $stock["Ticker"] . "</a></td>
                                        <td>" . $stock["OwnedStocks"] . "</td>
                                        <td>" . $stock["Group"] . "</td>
                                    </tr>
                                ";
                            }
                            ?>
                        </tbody>
                    </table>
                </div>

                <div class="card-body">
                    <form class="row g-2">
                        <div class="col-md-12">
                            <div class="input-group">
                                <span class="input-group-text">Ticker (e.g., AAPL, INTL)</span>
                                <input id="new-ticker-input" type="text" class="form-control">
                            </div>
                        </div>
                        <input id="new-ticker-btn" type='button' class='btn btn-success float-right' value='Track New Ticker'></input>
                    </form>
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
                    <h5 class="card-title">Research</h5>

                    <form class="row g-2">
                        <div class="col-md-12">
                            <div class="input-group">
                                <span class="input-group-text">Ticker (e.g., AAPL, INTC)</span>
                                <input id="ticker-search-input" type="text" class="form-control">
                            </div>
                        </div>
                        <input id="ticker-search-btn" type='button' class='btn btn-success float-right' value='Quick Search'></input>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">
  $(window).on("load", function()
  {
    $("#new-ticker-btn").click(function() 
    {
        var ticker = $("#new-ticker-input").val();
        console.log("Adding new ticker: " + ticker)

        $.ajax(
            {
            url: "<?php echo $GLOBALS['api']; ?>" + "/put/tracking/ticker",
            type: "POST",
            data: {
                "Ticker": ticker,
            },
            success: function (response) {
                location.reload(true);
            }
        });
    });

    $("#ticker-search-btn").click(function() 
    {
        var ticker = $("#ticker-search-input").val();
        window.location.replace("/tracking/stock.php?ticker=" + ticker);
    });
  });
</script>

<?php foot(); ?>