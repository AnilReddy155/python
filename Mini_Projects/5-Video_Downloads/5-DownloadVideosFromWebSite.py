import requests
import m3u8

import subprocess
import random


def download_video_series(video_links): 

 for link in video_links:

    # create response object  
    r = requests.get(link, stream = True) 
    m3u8_master = m3u8.loads(r.text)
    print(m3u8_master)
    data = m3u8_master.data
    playlist_url = m3u8_master.data['playlists']
    # r = requests.get(m3u8_master.data['playlists'][0]['uri'])
    # with open("vi.ts", 'wb') as f: 
    #    for chunk in r.iter_content(chunk_size = 1024*1024): 
    #        if chunk: 
     #          f.write(chunk) 

    file_name = "video-downloads/video"+str(random.randint(1,9))
    
    with open(file_name+".ts", 'wb') as f:
      for segment in m3u8_master.data['playlists']:
            url = segment['uri']
            r = requests.get(url)
            f.write(r.content)

    subprocess.run(['ffmpeg','-protocol_whitelist', 'file,http,https,tcp,tls,crypto' ,'-i', file_name+'.ts', '-c', 'copy', '-bsf:a','aac_adtstoasc' ,file_name+'.mp4'])
    
 

dwn_link = ['https://manifest.prod.boltdns.net/manifest/v1/hls/v4/clear/3588749423001/5c4ada4e-79d8-4824-b13e-1b91ac66910e/10s/master.m3u8?fastly_token=NjE0MGMzMjVfYzA3ZWI1OWQ2YTg0MjM5ODRmNjg3ODc0MDJhYWY4ZThmNTVmMTE4ZWNjZmRmMTNlN2NjNzNiYmI4ZDY1YzE5Yw%3D%3D']
# dwn_link = ['https://youtu.be/AfCaoIFLe6Q','https://youtu.be/2Z64G-Wyhqg']
# ffmpeg -protocol_whitelist file,http,https,tcp,tls,crypto -i MIE.m3u8 -c copy -bsf:a aac_adtstoasc MIE.mp4

download_video_series(dwn_link)
# subprocess.run(['ffmpeg', '-i', 'video1.ts', 'video1.mp4'])
# subprocess.run(['ffmpeg', '-i', 'video1.ts', 'video1.mp4'])

