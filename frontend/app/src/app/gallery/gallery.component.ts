import { Component, OnInit, Input } from '@angular/core';
import { ImageService } from '../_services/image.service'


@Component({
  selector: 'app-gallery',
  templateUrl: './gallery.component.html',
  styleUrls: ['./gallery.component.css']
})
export class GalleryComponent implements OnInit {
  images: string[];
  imageFile: File;
  @Input() isCurrentUser: boolean;
  @Input() routeID: string;
  picURL: any;
  isImageSaved: boolean;
  id: string;
  errorMsg: string;

  constructor(private imageService : ImageService) { }

  ngOnInit(): void {
    this.getImages();
    this.isImageSaved = false;
  }

  changeOutput() {
    if (document.getElementById("gallery").className == "square-gallery"){
      document.getElementById("gallery").className = "list-gallery";
    } else {
      document.getElementById("gallery").className = "square-gallery"; 
    } 
  }

  isGrid(): boolean {
    if (document.getElementById("gallery").className == "square-gallery"){
      return true;
    } else {
      return false; 
    } 
  } 

  addImage(): void {
    if (this.imageFile == null){
      return;
    }
    else{
      const uploadData = new FormData();
      uploadData.append('file', this.imageFile);
      this.imageService.uploadImage(uploadData).subscribe((out) => {
        this.images.unshift(out);
        this.isImageSaved = false;
      });
    }
  } 

  onFileChange(event) {
    this.errorMsg = "";
    this.imageFile = event.target.files[0];

    var mimeType = event.target.files[0].type;
		
		if (mimeType.match(/image\/*/) == null) {
      this.errorMsg = "Sorry, uploaded file must be an image"
			return;
    }
    
		var reader = new FileReader();
		reader.readAsDataURL(event.target.files[0]);
		
		reader.onload = (_event) => {
			this.picURL = reader.result; 
    }
    this.isImageSaved = true;
  }

  getImages(): void {
    this.imageService.getGallery(this.routeID).subscribe((images) => {
      if(images == null){
        this.images = [];
      }
      else {
        this.images = images;
      }
    });
  }

  removeImage() {
    this.imageFile = null;
    this.isImageSaved = false;
  }

  deleteImage() {
    console.log("Deletededed")
  }

  makeDeletable() {
    console.log("Deletededed")
  }
}
