import {expect, test} from '@playwright/test';

test('check all diamond planner routes', async ({ page }) => {
  await page.goto('/');
  await page.getByRole('link', { name: 'Login' }).click();
  await page.getByRole('textbox', { name: 'Your email' }).click();
  await page.getByRole('textbox', { name: 'Your email' }).fill('karltestmichel@obnx.dev');
  await page.getByRole('textbox', { name: 'Your password' }).click();
  await page.getByRole('textbox', { name: 'Your password' }).fill('testtest');
  await page.getByRole('button', { name: 'Login to your account' }).click();
  await page.getByRole('link', { name: 'Personal Stats' }).click();
  await page.getByRole('heading', { name: 'All Teams' }).click();
  await page.locator('section').filter({ hasText: 'Attendance Stats In 8% Maybe' }).locator('footer').click();
  await page.getByRole('link', { name: 'Player Profile' }).click();
  await page.getByRole('button', { name: 'Edit Player Data' }).click();
  await page.getByRole('button', { name: 'Close' }).click();
  await page.getByRole('link', { name: 'Fangzauner Schneebrunzer (ÖÄ' }).click();
  await page.getByRole('link', { name: 'Bogus Team (ÖÄZZ)' }).click();
  await page.getByRole('link', { name: 'Admin Dashboard' }).click();
  await expect(page.getByRole('heading', { name: 'Stats Admin Dashboard' })).toBeVisible();
  await page.getByRole('link', { name: 'Dashboard', exact: true }).click();
  await expect(page.getByRole('heading', { name: 'User Data' })).toBeVisible();
  await page.getByRole('link', { name: 'Bogus Team (ÖÄZZ)' }).click();
  await expect(page.getByRole('main')).toContainText('Dies ist ein Test-Team. Wenn es dir im realen Betrieb angezeigt wird, sag David Bescheid und schicke einen Screenshot mit.');
  await expect(page.getByText('Bogus Team Name 0 BSM-Liga (f')).toBeVisible();
  await page.locator('[id="radio-group\\:c19"] label').filter({ hasText: 'Past' }).click();
  await page.getByRole('link', { name: 'Player List' }).click();
  await expect(page.getByRole('heading', { name: 'Team members for Team "Bogus' })).toBeVisible();
});