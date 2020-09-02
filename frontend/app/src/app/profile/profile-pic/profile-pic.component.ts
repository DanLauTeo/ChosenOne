import { Component, OnInit } from '@angular/core';
import * as _ from 'lodash';
import { AccountService } from '../../_services/account.service';
import { User  } from '../../_models/user';

@Component({
  selector: 'app-profile-pic',
  templateUrl: './profile-pic.component.html',
  styleUrls: ['./profile-pic.component.css']
})
export class ProfilePicComponent implements OnInit {

  imageError: string;
  isImageSaved: boolean;
  cardImageBase64: string;
  user : User;

  constructor( private accountService : AccountService) {
    //this.accountService.user.subscribe(x => this.user = x);
    this.user = accountService.getUser();
  }

  ngOnInit() {
  }

  /*fileChangeEvent(fileInput: any) {
      this.imageError = null;
      if (fileInput.target.files && fileInput.target.files[0]) {
          // Size Filter Bytes
          const max_size = 20971520;
          const allowed_types = ['image/png', 'image/jpeg'];
          const max_height = 15200;
          const max_width = 25600;

          if (fileInput.target.files[0].size > max_size) {
              this.imageError =
                  'Maximum size allowed is ' + max_size / 1000 + 'Mb';

              return false;
          }

          if (!_.includes(allowed_types, fileInput.target.files[0].type)) {
              this.imageError = 'Only Images are allowed ( JPG | PNG )';
              return false;
          }

          const reader = new FileReader();
          reader.onload = (e: any) => {
              const image = new Image();
              image.src = e.target.result;
              image.onload = rs => {
                  const img_height = rs.currentTarget['height'];
                  const img_width = rs.currentTarget['width'];

                  if (img_height > max_height || img_width > max_width) {
                      this.imageError =
                          'Maximum dimensions allowed ' +
                          max_height +
                          'x' +
                          max_width +
                          'px';
                      return false;
                  } else {
                      const imgBase64Path = e.target.result;
                      this.cardImageBase64 = imgBase64Path;
                      this.isImageSaved = true;
                      // this.previewImagePath = imgBase64Path;
                  }
              };
          };

          reader.readAsDataURL(fileInput.target.files[0]);
      }
  }*/

  fileChangeEvent(fileInput: any) {
    this.imageError = null;
    if (fileInput.target.files && fileInput.target.files[0]) {
        // Size Filter Bytes
        const max_size = 20971520;
        const allowed_types = ['image/png', 'image/jpeg'];
        const max_height = 15200;
        const max_width = 25600;

        if (fileInput.target.files[0].size > max_size) {
            this.imageError =
                'Maximum size allowed is ' + max_size / 1000 + 'Mb';

            return false;
        }

        if (!_.includes(allowed_types, fileInput.target.files[0].type)) {
            this.imageError = 'Only Images are allowed ( JPG | PNG )';
            return false;
        }

        const reader = new FileReader();
        reader.onload = (e: any) => {
            const image = new Image();
            image.src = e.target.result;
            image.onload = rs => {
                const img_height = rs.currentTarget['height'];
                const img_width = rs.currentTarget['width'];

                if (img_height > max_height || img_width > max_width) {
                    this.imageError =
                        'Maximum dimensions allowed ' +
                        max_height +
                        'x' +
                        max_width +
                        'px';
                    return false;
                } else {
                    const imgBase64Path = e.target.result;
                    this.cardImageBase64 = imgBase64Path;
                    this.isImageSaved = true;
                    // this.previewImagePath = imgBase64Path;
                }
            };
        };

        reader.readAsDataURL(fileInput.target.files[0]);
    }
}

  removeImage() {
      this.cardImageBase64 = null;
      this.isImageSaved = false;
  }

}