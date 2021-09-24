import requests


video_download_folder = 'video-downloads/'
def download_file(urls):
    try:
        for url in urls:
            local_filename = url.split('/')[-1]
            r = requests.get(url, stream=True)
            
            with open(video_download_folder+local_filename, 'wb') as f:
                for chunk in r.iter_content(chunk_size=1024): 
                    if chunk: # filter out keep-alive new chunks
                        f.write(chunk)
        print("Videos Downloaded Successfully! ..")                
    except Exception as e:
        print(e)
        print("Error while downloading videos ") 

urls = ["http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
        "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4",
        "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerEscapes.mp4"]
download_file(urls)