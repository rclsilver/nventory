import { Component, Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-headers',
  templateUrl: './headers.component.html',
  styleUrls: ['./headers.component.scss'],
})
@Injectable({
  providedIn: 'root',
})
export class DebugHeadersComponent implements OnInit {
  public headers: Object;

  constructor(
    private http: HttpClient
  ) {}

  ngOnInit(): void {
    this.refresh();
  }

  getHeaders(): Observable<Object> {
    return this.http.get<Object>('/api/debug/headers');
  }

  refresh() {
    this.getHeaders().subscribe((headers: Object) => {
      this.headers = headers;
    });
  }
}
