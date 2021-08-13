# Amazon Product Scraper

This Lambda function scrapes an Amazon product page.

Input object:

```json
{
  "url": "https://www.amazon.com/Terrarium-Planter-Tabletop-Hydroponics-Decoration/dp/B07D29P5Z1/ref=sr_1_2?dchild=1&keywords=plant&qid=1628790382&sr=8-2"
}
```

Output object:

```json
{
  "sale_price": "$19.99",
  "original_price": "",
  "title": "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nRetro Wooden Frame Glass Plant Terrarium for Desktop Rustic wood and vintage design, these decorative glass vases are perfect for adding a touch of green and elegance to your home or office.  Packing included:  1 x Hexagon Screwdriver 2 x Screws 1 x Metal Swivel Holder 8 x Fixed Plug 1 x Wooden Stand ( 2 piece )  3 x Mini Bulb Shape Vase( No plant or other decorative objects included in this item.)  About Terrarium Size -The wooden stand size : 5.5\"H x 11\" W x 4\" D bulb mini vase: 3.74 H x 2.75 W, Opening – 1 inch Diameter Material -wooden frame, three high boron silicon heat resistant glass bulb shape vase\n\n\n\n\n\n\n\n\n\n\n",
  "description": ""
}
```

This needs the API Gateway to proxy an HTTP request to the Lambda function.

