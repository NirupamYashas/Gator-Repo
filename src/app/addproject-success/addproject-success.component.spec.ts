import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddprojectSuccessComponent } from './addproject-success.component';

describe('AddprojectSuccessComponent', () => {
  let component: AddprojectSuccessComponent;
  let fixture: ComponentFixture<AddprojectSuccessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddprojectSuccessComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddprojectSuccessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
