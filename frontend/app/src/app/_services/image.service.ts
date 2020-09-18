import { HttpClient } from  '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ImageService {

  constructor(private httpClient: HttpClient) { }

  public getFeed() {
    return this.httpClient.get<any>(`/feed-images/`);
  }

  public getGallery(id: string) {
    return this.httpClient.get<any>(`/user/${id}/images/`);
  }

  public uploadImage(imageData: FormData) {
    return this.httpClient.post<any>(`/images/`, imageData);
  }

  public deleteImage(imgID: number) {
    return this.httpClient.delete<any>(`/images/${imgID}/`);
  }
}
