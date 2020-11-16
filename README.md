# WeatherDiffService
 
A simple weather service allows observing differences between weather data suppliers.


The standard procedure allows the GET method only. Requests which use other methods will be denied. To use WeatherDiffService, the request should be properly combined. The service expects two incoming parameters "country" and "city".

"country" parameter should contain two symbols "GB" or "us".

"city" is the city name like "london" or "Moscow"

All parameters are case insensitive.


E.g.
http://example.com/?country=us&city=sain-perersburg
http://example.com/?country=GB&city=BRISTOL