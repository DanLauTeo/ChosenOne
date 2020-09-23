import { Component, OnInit, Input } from '@angular/core';
import { ImageService } from '../_services/image.service'
import { GalleryImage } from '../_models/galleryImage'
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { PopupComponent } from './popup/popup.component';

@Component({
  selector: 'app-gallery',
  templateUrl: './gallery.component.html',
  styleUrls: ['./gallery.component.css']
})
export class GalleryComponent implements OnInit {
  images: GalleryImage[];
  imageFile: File;
  @Input() isCurrentUser: boolean;
  @Input() routeID: string;
  picURL: any;
  isImageSaved: boolean;
  id: string;
  errorMsg: string;
  canDelete: boolean;

  constructor(private imageService : ImageService, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getImages();
    this.isImageSaved = false;
    this.canDelete = false;
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
        console.log(images);
        this.images = images;
      }
    });
  }

  removeImage() {
    this.imageFile = null;
    this.isImageSaved = false;
  }

  cancel() {
    this.canDelete = false;
  }

  allowDelete() {
    this.canDelete = true;
  }

  deleteImage(image: GalleryImage) {
    const dialogRef = this.dialog.open(PopupComponent, {
      width: '250px',
      data: image
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result != null){
        this.imageService.deleteImage(result).subscribe((res) => {
          console.log(res);
          this.images = this.images.filter(({ imgID }) => imgID !== result); 
          this.canDelete = false;
        });       
      }
      else{
        console.log("Hmmm");
      }
    });
  }
}
