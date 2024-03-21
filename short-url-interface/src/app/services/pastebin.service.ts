import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Result, ShortURL, Stats, URLData } from '../models/model';

@Injectable({
  providedIn: 'root'})
export class PastebinService {

  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  showAll(): Observable<Result> {
    return this.http.get<Result>(`${this.apiUrl}/`);
  }

  redirect(shorturl: string): Observable<void> {
    return this.http.get<void>(`${this.apiUrl}/${shorturl}`);
  }

  shorten(urlData: URLData): Observable<ShortURL> {
    return this.http.post<ShortURL>(`${this.apiUrl}/`, urlData);
  }

  stats(shorturl: string): Observable<Stats> {
    return this.http.get<any>(`${this.apiUrl}/${shorturl}/stats`);
  }
}