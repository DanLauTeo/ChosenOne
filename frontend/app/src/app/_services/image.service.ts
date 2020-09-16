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
}
