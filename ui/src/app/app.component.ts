import { Component, OnInit, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
@Injectable({
  providedIn: 'root',
})
export class AppComponent implements OnInit {
  title = 'nventory';
  user = null;

  constructor(
    private http: HttpClient
  ) {}

  ngOnInit(): void {
    this.http.get<Object>('/api/auth/user').subscribe((user: Object) => {
      this.user = user;
    });
  }
}
