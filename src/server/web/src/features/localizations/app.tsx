import React, { useState } from 'react';
import { Select, Option, Typography, Box, CircularProgress } from '@mui/joy';
import { useGetLocalizationsQuery } from './localizationsSlice'; // Import the hook and interface

const Localizations: React.FC = () => {
  const { data: localData, error, isLoading: loading } = useGetLocalizationsQuery(); // Use the hook
  const [selectedLanguage, setSelectedLanguageState] = useState<string | null>('en-US');
  const [selectedCurrency, setSelectedCurrencyState] = useState<string | null>('USD');

  const handleLanguageChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedLanguageState(newValue); // Using local state for now
      // Removed logic to auto-select currency based on language
    }
  };

  const handleCurrencyChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedCurrencyState(newValue); // Using local state for now
    }
  };

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return <Typography color="danger">Error fetching data: {('data' in error ? JSON.stringify(error.data) : "error.message") || 'Unknown error'}</Typography>;
  }

  if (!localData) {
    return <Typography>No data available.</Typography>;
  }

  const allLanguages = localData.flatMap(entry =>
    entry.Localizations ? entry.Localizations.map(loc => loc.Language) : []
  );
  const uniqueLanguages = Array.from(
    new Map(
      allLanguages
        .filter(lang => lang && lang.ID != null && lang.LanguageCulture != null && lang.Name != null)
        .map(lang => [lang.ID, lang])
    ).values()
  );

  const allCurrencies = localData.flatMap(entry =>
    entry.Localizations ? entry.Localizations.map(loc => loc.Currency) : []
  );
  const uniqueCurrencies = Array.from(
    new Map(
      allCurrencies
        .filter(curr => curr && curr.ID != null && curr.CurrencyCode != null && curr.Name != null)
        .map(curr => [curr.ID, curr])
    ).values()
  );

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, p: 2 }}>
      <Typography level="h4" component="h1">
        Localization Settings
      </Typography>

      <Box>
        <Typography gutterBottom>Language</Typography>
        <Select
          value={selectedLanguage}
          onChange={handleLanguageChange}
          placeholder="Select language"
        >
          {uniqueLanguages.map((lang) => (
            <Option key={lang.ID} value={lang.LanguageCulture}>
              {lang.Name} ({lang.LanguageCulture})
            </Option>
          ))}
        </Select>
      </Box>

      <Box>
        <Typography gutterBottom>Currency</Typography>
        <Select
          value={selectedCurrency}
          onChange={handleCurrencyChange}
          placeholder="Select currency"
        >
          {uniqueCurrencies.map((currency) => (
            <Option key={currency.ID} value={currency.CurrencyCode}>
              {currency.Name} ({currency.CurrencyCode})
            </Option>
          ))}
        </Select>
      </Box>
       {
        /*
        <Box sx={{ mt: 2 }}>
        <Typography>Selected Language: {selectedLanguage || 'None'}</Typography>
        <Typography>Selected Currency: {selectedCurrency || 'None'}</Typography>
      </Box>
        */
       }   
      
    </Box>
  );
};

export default Localizations;
