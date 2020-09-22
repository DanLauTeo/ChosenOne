import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import * as _ from 'lodash';
import { AccountService } from '../../_services/account.service';
import { User  } from '../../_models/user';
import { ProfileService } from '../../_services/profile.service'

@Component({
  selector: 'app-profile-pic',
  templateUrl: './profile-pic.component.html',
  styleUrls: ['./profile-pic.component.css']
})
export class ProfilePicComponent implements OnInit {
  @Input() id: string;
  @Input() picURL: any;
  @Input() isEditable: boolean;
  currentPic: any;
  imageFile: File;
  isImageSaved: boolean;
  @Output() userOut = new EventEmitter<User>(); 

  constructor(private profileService : ProfileService) {
  }

  ngOnInit() {
    this.imageFile = null;
    this.isImageSaved = false;
    this.currentPic = this.picURL;
  }

  onFileChange(event) {
    this.imageFile = event.target.files[0]

    this.isImageSaved = true;

    var mimeType = event.target.files[0].type;
		
		if (mimeType.match(/image\/*/) == null) {
			return;
		}
		
		var reader = new FileReader();
		reader.readAsDataURL(event.target.files[0]);
		
		reader.onload = (_event) => {
			this.currentPic = reader.result; 
		}
  }

  onUpload() {
    if (this.imageFile == null){
      return;
    }
    else{
      const uploadData = new FormData();
      uploadData.append('file', this.imageFile);
      this.profileService.uploadProfilePic(this.id, uploadData).subscribe((user:User) => {
        this.picURL = user.profilePic;
        this.userOut.emit(user);
        this.isImageSaved = false;
      });
    }
  }

  removeImage() {
    this.currentPic = this.picURL;   
    this.isImageSaved = false;
    this.imageFile = null;
  }

}
