import { HttpClient } from  '@angular/common/http';  
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  routeURL: string ;

  constructor(private httpClient: HttpClient) { }

  public getUser(id) {
    this.routeURL = "user/"+id;
    return this.httpClient.get<any>(this.routeURL);
  }

  public patchUser(id, patch) {
    this.routeURL = "user/"+id;
    return this.httpClient.patch<any>(this.routeURL, patch);
  }
  
  public uploadProfilePic(id, formData) {
    this.routeURL = "user/"+id+"/profile-image";
    return this.httpClient.put<any>(this.routeURL, formData);
  }
}
