import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-gallery',
  templateUrl: './gallery.component.html',
  styleUrls: ['./gallery.component.css']
})
export class GalleryComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  changeOutput() {
    if (document.getElementById("gallery").className == "square-gallery"){
      document.getElementById("gallery").className = "list-gallery";
      document.getElementById("gridList").innerText = "Grid";
    } else {
      document.getElementById("gallery").className = "square-gallery";
      document.getElementById("gridList").innerText = "List";      
    } 
  }
}
