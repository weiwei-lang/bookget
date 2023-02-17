# bookget
bookget 数字图书馆下载工具

支持的数字图书馆URL格式，请注意：一般以可以在线阅读的网址为准。
### 中国地区数字图书馆：
1. [中国][国家图书馆](http://read.nlc.cn/thematDataSearch/toGujiIndex)
```
整书多册URL：http://read.nlc.cn/allSearch/searchDetail?searchType=1002&showType=1&indexName=data_892&fid=411999021002
或者单册URL：http://read.nlc.cn/OutOpenBook/OpenObjectBook?aid=403&bid=70621.0
```

2. [中国][香港中文大学图书馆](https://repository.lib.cuhk.edu.hk/sc/collection)
```
https://repository.lib.cuhk.edu.hk/sc/item/cuhk-412225#page/1/mode/2up
```

3. [中国][香港科技大学图书馆](https://lbezone.ust.hk/)
```
https://lbezone.ust.hk/bib/b1129168
```

4. [中国][台北古籍与特藏文献](https://rbook.ncl.edu.tw/NCLSearch)
```
https://rbook.ncl.edu.tw/NCLSearch/Search/SearchDetail?item=422a7598bd0046aebf2684ae0f945d25fDcyODIz0&image=1&page=&whereString=&sourceWhereString=&SourceID=
```

5. [中国][台北故宫博物院 – 善本古籍 ](https://rbk-doc.npm.edu.tw/)   
   注：查阅[参考PDF文档](/doc/pdf/05.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E5%8F%B0%E5%8C%97%E6%95%85%E5%AE%AB%E5%8D%9A%E7%89%A9%E9%99%A2%E5%96%84%E6%9C%AC%E5%8F%A4%E7%B1%8D.pdf)

6. [中国][四川古籍数字图书馆](http://guji.sclib.org/qt-zxsk.html)   
   注：需配合 dezoomify-rs zhudw 修改版，才可以下载。
```
http://guji.sclib.org/qt-tsxq.html?id=52
http://guji.sclib.org/viewer.html?bookId=52#page=18&viewer=picture
```

7. [中国][云南古籍数字图书馆](http://msq.ynlib.cn/)   
   注：需配合 dezoomify-rs zhudw 修改版，才可以下载。
```
http://msq.ynlib.cn/#/details/?id=1001
http://msq.ynlib.cn/frontend/viewer.html?typeId=80&bookId=1001#/page=1&viewer=picture
```

8. [中国][天一阁博物院古籍数字图书馆](https://gj.tianyige.com.cn/)
   注：此站点需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
https://gj.tianyige.com.cn/SearchPage/c56c5afbb95f667c96c57b6d3b4c5f0c
https://gj.tianyige.com.cn/Book?catalogId=c56c5afbb95f667c96c57b6d3b4c5f0c&directoryId=5adf0ab25f361ec4c96205023079c8b9&fascicleId=c3e3ee09cfbb2059c586207344310943
```

9. [中国][广州大典](http://gzdd.gzlib.gov.cn/Hrcanton/)   
   注：可通过微信小程序【粤通读】注册读者证。此站点需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultDetail?BookId=GZDD022601004
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultSummary?bookid=GZDD022601004&filename=GZDD022601004#
```

10. [中国][深圳市古籍数字图书馆](https://yun.szlib.org.cn/stgj2021/)
```
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=4&book_id=18269  
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=1&book_id=18017
```

11. [中国][洛阳市图书馆](http://221.13.137.120:8090/index.php)
```
http://221.13.137.120:8090/productshow.php?cid=4&id=112
```

12. [中国][温州市图书馆-瓯越记忆](https://oyjy.wzlib.cn/pdf/)
```
https://oyjy.wzlib.cn/resource/?id=61e4c764505415b2e6921e5e
https://oyjy.wzlib.cn/resource/?id=62c56bb357de1ef36b1f5614
```

### 欧美数字图书馆：
1. [美国][哈佛大学图书馆](https://hollis.harvard.edu/) [或哈佛燕京图书馆藏](https://gj.library.sh.cn/org/harvard)   
```
https://iiif.lib.harvard.edu/manifests/view/drs:53262215
```

2. [美国][hathitrust数字图书馆](https://www.hathitrust.org/)
```
https://babel.hathitrust.org/cgi/pt?id=uc1.c087423515&view=1up&seq=1&skin=2021
```

3. [美国][普林斯顿大学图书馆](https://library.princeton.edu/)
```
https://catalog.princeton.edu/catalog/9940468523506421
https://dpul.princeton.edu/catalog/99915a8b423b596e47540e3feeee19b8
```

4. [美国][国会图书馆](https://www.loc.gov/collections/chinese-rare-books/)   
   注：中国大陆访问此网站需自备海外VPN，免VPN方法需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
https://www.loc.gov/item/2014514163/
```

5. [美国][斯坦福大学图书馆](https://searchworks.stanford.edu/?f%5Baccess_facet%5D%5B%5D=Online&f%5Bbuilding_facet%5D%5B%5D=East+Asia&f%5Bformat_main_ssim%5D%5B%5D=Book&f%5Blanguage%5D%5B%5D=Chinese&utf8=%E2%9C%93)
```
https://searchworks.stanford.edu/view/4182111   
```

6. [美国][familysearch.org 中國族譜收藏 1239-2014年](https://www.familysearch.org/search/collection/1787988)   
   注：此站点需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-C8KN?wc=3X27-MNY%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022419701&cc=1787988
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-CRG9?wc=3X2Q-BZ7%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022421801&cc=1787988
```
- [美国][familysearch.org 家譜圖像](https://www.familysearch.org/records/images/)
```
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9S9-WS92?view=explore&groupId=M94X-6HR
https://www.familysearch.org/records/images/image-details?rmsId=M94F-78D&jiapuOnly=true&surname=%E6%9C%B1&place=2013&showUnknown=true&page=1&pageSize=100&imageIndex=0
```

7. [德国][柏林国立图书馆](https://digital.staatsbibliothek-berlin.de)
```
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770&PHYSID=PHYS_0001
```

8. [德国][巴伐利亚州立图书馆](https://ostasien.digitale-sammlungen.de/)
```
https://ostasien.digitale-sammlungen.de/view/bsb11129280/1
```
9. [英国][牛津大学博德利图书馆](https://digital.bodleian.ox.ac.uk/collections/chinese-digitization-project/)
```
https://digital.bodleian.ox.ac.uk/objects/310cb04e-6bce-44e3-85b5-03417c9644a8/
```
10. [英国][大英图书馆（手稿本）](http://www.bl.uk/manuscripts/)
```
http://www.bl.uk/manuscripts/Viewer.aspx?ref=or_6814!1_fs001r
```



### 日本数字图书馆：

1. [日本][京都大学图书馆](https://rmda.kulib.kyoto-u.ac.jp/)

```
https://rmda.kulib.kyoto-u.ac.jp/item/rb00024956
```
2. [日本][国立国会图书馆](http://dl.ndl.go.jp/)
```
https://dl.ndl.go.jp/info:ndljp/pid/8929985
```
3. [日本][E国宝e-Museum]( https://emuseum.nich.go.jp/)
```
https://emuseum.nich.go.jp/detail?content_base_id=100168&content_part_id=009&langId=zh&webView=
```
4. [日本][宫内厅书陵部](https://db2.sido.keio.ac.jp/kanseki/T_bib_search.php)
```
https://db2.sido.keio.ac.jp/kanseki/T_bib_frame.php?id=006754
```
5. [日本][东京大学东洋文化研究所](http://shanben.ioc.u-tokyo.ac.jp/list.php)
```
http://shanben.ioc.u-tokyo.ac.jp/main_p.php?nu=C5613401&order=rn_no&no=00870
```

6. [日本][国立公文书馆（内库文库）](https://www.digital.archives.go.jp/)
```
https://www.digital.archives.go.jp/DAS/meta/listPhoto?LANG=default&BID=F1000000000000095447&ID=&NO=&TYPE=
```
7 [日本][东洋文库]( http://dsr.nii.ac.jp/toyobunko/index.html.ja)
```
http://dsr.nii.ac.jp/toyobunko/XI-6-A-16/V-1/
```
8. [日本][早稻田大学图书馆](https://www.wul.waseda.ac.jp/kotenseki/search.php)
```
https://archive.wul.waseda.ac.jp/kosho/ri08/ri08_01899/
```
9. [日本][新日本古典籍综合数据库](https://kotenseki.nijl.ac.jp/)   
   注：查阅[参考pdf文档](/doc/pdf/04.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E6%96%B0%E6%97%A5%E6%9C%AC%E5%8F%A4%E5%85%B8%E5%9B%BE%E4%B9%A6.pdf)
```
https://kotenseki.nijl.ac.jp/biblio/100270332/viewer/1
https://kotenseki.nijl.ac.jp/biblio/100270332
```

10. [日本][京都大学人文科学研究所 - 东方学数字图书博物馆](http://kanji.zinbun.kyoto-u.ac.jp/db-machine/toho/html/top.html)
```
http://kanji.zinbun.kyoto-u.ac.jp/db-machine/toho/ShiSanJingZhuShu/html/A002menu.html
```

11. [日本][国立历史民俗博物馆](https://khirin-a.rekihaku.ac.jp/)
```
单册URL:
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-1
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-173-1

多册URL，使用和“批量下载”相同格式，但是无需修改config.ini中配置。
如：第1-9册，第10-90册。用圆括号包围数字。   

https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-(1-90)
https://khirin-a.rekihaku.ac.jp/sohankanjo/h-173-(1-61)
```
12. [日本][市立米泽图书馆](https://www.library.yonezawa.yamagata.jp/dg/zen.html)
```
https://www.library.yonezawa.yamagata.jp/dg/AA001_view.html
https://www.library.yonezawa.yamagata.jp/dg/AA002_view.html
```
13. [日本][庆应义塾大学图书馆](https://dcollections.lib.keio.ac.jp/ja/kanseki)
```
https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1
```
14. [日本][关西大学图书馆](https://www.iiif.ku-orcas.kansai-u.ac.jp/books)
```
https://www.iiif.ku-orcas.kansai-u.ac.jp/books/210185040#?page=1
```


### 其它数字图书馆：
1. [世界][國際敦煌項目](http://idp.nlc.cn/)   
   注：需先搜索关键词，例如`8210`，并且URL中含有`uid=xxxx`，短时间内有效，请在搜索结果后尽快下载。
```
http://idp.nlc.cn/database/oo_scroll_h.a4d?uid=47355195088;recnum=0;index=2
```
2.  [韩国][国家图书馆](https://www.dlibrary.go.kr/) [或开放数据](https://lod.nl.go.kr/)
   注：请使用v0.2.6版。新版不再支持。查阅[参考pdf文档](/doc/pdf/03.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E9%9F%A9%E5%9B%BD%E5%9B%BE%E4%B9%A6%E9%A6%86%E5%9B%BE%E4%B9%A6.pdf)
```
http://lod.nl.go.kr/page/CNTS-00076977176
```

