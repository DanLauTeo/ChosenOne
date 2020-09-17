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
  imageFile: File;
  isImageSaved: boolean;
  user : User;
  @Output() userOut = new EventEmitter<User>(); 

  constructor( private accountService : AccountService, private profileService : ProfileService) {
    //this.accountService.user.subscribe(x => this.user = x);
    this.user = accountService.getUser();
  }

  ngOnInit() {
    this.imageFile = null;
    this.isImageSaved = false;
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
			this.picURL = reader.result; 
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
        this.userOut.emit(user);
      });
    }
    
  }

  removeImage() {
    this.imageFile = null;
    this.isImageSaved = false;
    this.picURL = this.user.profilePic;
  }

}
