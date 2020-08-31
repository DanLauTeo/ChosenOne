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
      //document.getElementById("img").src="../../assets/icons/view_module-24px.svg";
    } else {
      document.getElementById("gallery").className = "square-gallery";
      //document.getElementById("img").src="../../assets/icons/check_box_outline_blank-24px.svg";   
    } 
  }
}
