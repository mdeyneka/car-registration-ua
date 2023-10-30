#!/usr/bin/env python3

import argparse
import csv, sys
from datetime import datetime

parser = argparse.ArgumentParser(description='Transform data for car-registration-ua microservice',
                                 epilog='Data located at https://data.gov.ua/dataset/06779371-308f-42d7-895e-5a39833375f0')
parser.add_argument('--input', required=True, metavar='FILENAME', help='input file for transformation')
parser.add_argument('--output', default='output.csv', metavar='FILENAME', help='output file after transformation (default output.csv)')

def main():
    args = parser.parse_args()
    lines = []
    fieldnames = ['PERSON', 'REG_ADDR_KOATUU', 'OPER_CODE', 'OPER_NAME', 'D_REG', 'DEP_CODE', 'DEP', 'BRAND', 'MODEL', 'VIN', 'MAKE_YEAR', 'COLOR', 'KIND', 'BODY', 'PURPOSE', 'FUEL', 'CAPACITY', 'OWN_WEIGHT', 'TOTAL_WEIGHT', 'N_REG_NEW']
    with open(args.input, newline='') as csvfile:
        reader = csv.DictReader(csvfile, delimiter=';')
        reader.fieldnames = [name.upper() for name in reader.fieldnames]
        try:
            for row in reader:
                if ',' in row['OWN_WEIGHT']:
                    row['OWN_WEIGHT'] = row['OWN_WEIGHT'].replace(',', '.')
                if ',' in row['TOTAL_WEIGHT']:
                    row['TOTAL_WEIGHT'] = row['TOTAL_WEIGHT'].replace(',', '.')
                if row['FUEL'] == 'NULL':
                    row['FUEL'] = ''
                if row['CAPACITY'] == 'NULL':
                    row['CAPACITY'] = ''
                if row['OWN_WEIGHT'] == 'NULL':
                    row['OWN_WEIGHT'] = ''
                if row['TOTAL_WEIGHT'] == 'NULL':
                    row['TOTAL_WEIGHT'] = ''
                if row['N_REG_NEW'] == 'NULL':
                    row['N_REG_NEW'] = ''
                try:
                    date = datetime.strptime(row['D_REG'], '%Y-%m-%d')
                    row['D_REG'] = date.strftime('%d.%m.%Y')
                except ValueError:
                    pass
                lines.append(row)
        except csv.Error as e:
            sys.exit('file {}, line {}: {}'.format(args.input, reader.line_num, e))
    
    with open(args.output, 'w', newline='') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()
        for row in lines:
            writer.writerow(row)

if __name__ == "__main__":
    main()