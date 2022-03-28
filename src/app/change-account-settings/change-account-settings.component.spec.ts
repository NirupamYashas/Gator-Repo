import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeAccountSettingsComponent } from './change-account-settings.component';

describe('ChangeAccountSettingsComponent', () => {
  let component: ChangeAccountSettingsComponent;
  let fixture: ComponentFixture<ChangeAccountSettingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangeAccountSettingsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangeAccountSettingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
