import { HttpClient } from  '@angular/common/http';
import { map } from  'rxjs/operators';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MatchesService {

  constructor(private httpClient: HttpClient) { }

  public getMatches(): Observable<string[]> {
    return this.httpClient.get<MatchesResponse>("/matches/")
      .pipe(map(resp => resp.user_ids))
  }
}

class MatchesResponse {
    user_ids: string[];
}
