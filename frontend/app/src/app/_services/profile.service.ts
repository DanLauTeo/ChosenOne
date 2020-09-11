import { HttpClient, HttpEvent, HttpErrorResponse, HttpEventType, HttpParams } from  '@angular/common/http';  
import { map } from  'rxjs/operators';
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
    
  return this.httpClient.put(this.routeURL, formData);
  }
}
