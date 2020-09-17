import { HttpClient } from  '@angular/common/http';  
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ImageService {
  routeURL: string;

  constructor(private httpClient: HttpClient) { }

  public getFeed() {
    this.routeURL = "feed-images";
    return this.httpClient.get<any>(this.routeURL);
  }

  public getGallery(id) {
    this.routeURL = "user/"+id+"/images";
    return this.httpClient.get<any>(this.routeURL);
  }

  public uploadImage(imageData) {
    this.routeURL = "image-uploaded/";
    return this.httpClient.post<any>(this.routeURL, imageData);
  }

  public deleteImage(imgID) {
    this.routeURL = "images/"+imgID;
    return this.httpClient.delete<any>(this.routeURL);
  }
}
