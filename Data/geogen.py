import numpy as np
import numpy.random as nr
from datetime import datetime
import random
import json
import csv


#shanghai
# SJTU xuhui: 121.442393, 31.205193
# FDU: 121.509937, 31.304613

# minhang sports park: 121.376893, 31.151117
# xujiahui: 121.446395, 31.201964
# center: beiling dasha 121.413446, 31.178717
# size: array([ 0.065898,  0.046494])
# center as beiling, right-top xujiahui

# center: array([ 121.476165,   31.254903])
# size: array([ 0.067544,  0.09942 ])

# `count` gaussian points in a rect with `center`, `width`, `height`
def gaussian(count, center, size):
	center = np.array(center)
	size = np.array(size)
	points = nr.randn(count,2)
	# size = 3*sigma, or sigma = size / 3
	#fixme: 0.03% points would be abnormally on border
	points = np.clip(points, [-3, -3], [3,3])
	return points * (size/3) + center

def shpoints(count, center_count):
	center = np.array([ 121.413446, 31.178717])
	size = np.array([ 0.065898,  0.046494])

	minp = center - size / 2
	maxp = center + size / 2
	centers = np.array([nr.uniform(minp[0],maxp[0],center_count),nr.uniform(minp[1],maxp[1],center_count)]).T
	#fixme: I don't know how long is 0.1 deg
	points = [gaussian(count / center_count, center, [0.005,0.005]) for center in centers]
	return points
	# return np.reshape(points, (count, 2))
def shdata_notext(count, center_count):
	now = 1476539612
	yesterday = 1476439612
	points = shpoints(count, center_count)
	for i in range(center_count):
		single_distri = points[i]
		for point in single_distri:
			yield "{} {} {} {}\n".format(point[0],point[1],i % 10,random.randint(yesterday, now))


def shdata(count, center_count):
	items = []
	points = np.reshape(shpoints(count, center_count),(count, 2))

	with open('week1.csv') as csvfile:
		reader = csv.DictReader(csvfile)

		for i in range(count):
			row = reader.next()
			# 2012-01-03 01:12:55
			date = datetime.strptime(row['created_at'], '%Y-%m-%d %H:%M:%S')
			timestamp = (date - datetime(1970, 1, 1)).total_seconds()
			point = points[i]

			item = {}
			item['date'] = timestamp
			item['type'] = 'text'
			item['content'] = row['text']
			item['source'] = 'http://weibo.cn/some/path'
			item['location'] = {
				'lat': point[0],
				'lng': point[1]
			}
			items.append(item)
	return items

# import matplotlib.pyplot as plt
def main():

	# f = open('out.data','w')
	# for x in shdata_notext(50000,200):
	# 	print(x)
	# 	f.write(x)
	# f.close()
	f = csv.reader(open('out.data'), delimiter=' ')

	m = min([int(r[3]) for r in f])
	print(m)

	# data = shpoints(5000,200)
	# # print(json.dumps(data))
	# # ps = shdata(10000,100)
	# ps = np.reshape(data,(5000,2))
	# plt.scatter(ps[:,0],ps[:,1])
	# plt.show()

main()
