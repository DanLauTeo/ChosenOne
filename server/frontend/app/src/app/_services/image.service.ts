import { HttpClient } from  '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ImageService {

  basePath: string = "/api/v1"

  constructor(private httpClient: HttpClient) { }

  public getFeed() {
    return this.httpClient.get<any>(`${this.basePath}/feed-images/`);
  }

  public getGallery(id: string) {
    return this.httpClient.get<any>(`${this.basePath}/user/${id}/images/`);
  }

  public uploadImage(imageData: FormData) {
    return this.httpClient.post<any>(`${this.basePath}/images/`, imageData);
  }

  public deleteImage(imgID: number) {
    return this.httpClient.delete<any>(`${this.basePath}/images/${imgID}/`);
  }
}
