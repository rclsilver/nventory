import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DebugHeadersComponent } from './debug/headers/headers.component';

const routes: Routes = [
  { path: 'debug/headers', component: DebugHeadersComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
