import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
// import * as data from './combined2-copy.geojson'

import * as d3 from 'd3';

class App extends Component {
  drawMaps(json){
    console.log('json', json)
  }

  componentDidMount(){
    var projection = d3.geoConicConformal()
      .parallels([40 + 40 / 60, 41 + 2 / 60])
      .rotate([74, 0]);
      // .fitExtent([[20, 20], [940, 480]], nj);

    	//パスジェネレーター生成
    	// var path = d3.geoPath().projection(projection);　
      //
    	// //地図用のステージ(SVGタグ)を作成
    	// var map = d3.select("body")
    	// 	.append("svg")
    	// 	.attr("width", 960)
    		// .attr("height", 500);
        // console.log(data)

    	//地理データ読み込み
      d3.json("http://localhost:8080/gtrain", response => response.json())
      .then(data => console.log(data))

    	d3.json("/combined2-copy.geojson", response => response.json())
      .then(data => drawMaps(data))

    	//地図を描画
    	function drawMaps(geojson) {
        projection.fitExtent([[2, 2], [1920, 1000]], geojson);
        var path = d3.geoPath().projection(projection);　

        //地図用のステージ(SVGタグ)を作成
        var map = d3.select("body")
          .append("svg")
          .attr("width", 1920)
          .attr("height", 1000);
        console.log('geojson', geojson)
    		map.selectAll("path")
    			.data(geojson.features)
    			.enter()
    			.append("path")
    			.attr("d", path)  //パスジェネレーターを使ってd属性の値を生成している
    			.attr("fill", "green")
    			.attr("fill-opacity", 0.5)
    			.attr("stroke", "#222");
    	}

  }
  render() {
    return (
      <div className="App">
      </div>
    );
  }
}

export default App;
