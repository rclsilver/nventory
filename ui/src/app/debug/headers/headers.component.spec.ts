import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DebugHeadersComponent } from './headers.component';

describe('DebugHeadersComponent', () => {
  let component: DebugHeadersComponent;
  let fixture: ComponentFixture<DebugHeadersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DebugHeadersComponent ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DebugHeadersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
